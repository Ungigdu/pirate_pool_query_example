pragma solidity >=0.5.11;

import "./owned.sol";
import "./safemath.sol";
import "./IERC20.sol";
import "./Entity.sol";

contract PacketMarket is owned{

    event PoolRegister(
        address indexed pool,
        uint256 gtn);

    event PoolDestroy(
        address indexed pool,
        uint256 time);

    event PoolClaim(
        address indexed pool,
        address indexed user,
        uint256 packet,
        uint256 tonken,
        uint256 micrNonce,
        uint256 claimNonce);

    event UserRecharge(
        address indexed from,
        address indexed to,
        uint256 tokens,
        uint256 packets);

    event UserRefund(
        address indexed user,
        address indexed pool,
        uint256 tokens,
        uint256 packets);

    using SafeMath for uint256;
    using Entity for Entity.UserDetails;
    using Entity for Entity.MinerPool;
    using Entity for Entity.VersionManager;
    using Entity for Entity.SystemSetting;
    /******************
    * system settings
    ******************/

    IERC20 public token;
    uint public decimal = 18;
    address[] public PoolsIndex;
    Entity.VersionManager public vm =  Entity.VersionManager(0,0,0,1);
    Entity.SystemSetting public setting;


    mapping(address=>address[]) public poolsUsedByUser;
    mapping(address=>address[]) public usersUnderPool;
    mapping(address=>bytes32[]) public minersUnderPool;
    mapping(address=>Entity.MinerPool) public PoolData;
    mapping(bytes32=>Entity.Miner) public MinerData;
    mapping(address=>mapping(address=>Entity.UserDetails)) public UserData;

    constructor(address ta, uint d) public{
        require(ta != address(0) && d > 0);

        token = IERC20(ta);
        decimal = d;

        setting.MBytesPerToken = 1000;
        setting.Duration = 30 days;
        setting.MinPoolGuarantee = 102400 * (10 ** decimal);
        setting.MinMinerGuarantee = 50000 * (10 ** decimal);
    }

    function emergency() public onlyOwner{
        token.transfer(msg.sender, token.balanceOf(address(this)));
    }

    function changeSetting(uint price, uint duration, uint pGTN, uint mGTN) public onlyOwner{
        bool changed = false;
        if (price != 0){
            changed = true;
            setting.MBytesPerToken = price;
        }
        if (duration != 0){
            changed = true;
            setting.Duration = duration * 1 days;
        }
        if (pGTN != 0){
            changed = true;
            setting.MinPoolGuarantee = pGTN * (10 ** decimal);
        }
        if (mGTN != 0){
            changed = true;
            setting.MinMinerGuarantee = mGTN * (10 ** decimal);
        }

        if (changed){
            vm.upgradeVer(Entity.ActionType.SysSetting);
        }
    }

    function TokenBalance(address userAddress) public view returns (uint, uint, uint){
        return (token.balanceOf(userAddress), userAddress.balance, token.allowance(userAddress, address(this)));
    }

    /******************
    * pool action
    ******************/
    function RegPool(uint tokenNo, address poolAddr, bytes memory name, bytes memory email, bytes memory url) public{
        require(PoolData[poolAddr].mainAddr == address(0));
        require(tokenNo >= setting.MinPoolGuarantee);

        token.transferFrom(msg.sender, address(this), tokenNo);

        Entity.MinerPool memory pool = Entity.MinerPool(PoolsIndex.length, poolAddr, msg.sender, tokenNo, name, email, url);
        PoolsIndex.push(pool.mainAddr);

        PoolData[pool.mainAddr] = pool;
        vm.upgradeVer(Entity.ActionType.PoolReg);
        emit PoolRegister(poolAddr, tokenNo);
    }

    function DestroyPool(address poolAddr) public{
        Entity.MinerPool memory pool = PoolData[poolAddr];
        require(pool.payerAddr == msg.sender);

        delete PoolsIndex[pool.ID];
        delete PoolData[poolAddr];

        token.transfer(msg.sender, pool.GTN);

        vm.upgradeVer(Entity.ActionType.PoolExit);
        emit PoolDestroy(poolAddr, now);
    }

    function Pools() public view returns (address[] memory) {
        return PoolsIndex;
    }

    function usersOfPool(address pool) public view returns(address[] memory){
        return usersUnderPool[pool];
    }

    function ChangeBasicInfo(address mainAddr, bytes memory name, bytes memory email, bytes memory url) public{
        Entity.MinerPool storage pool = PoolData[mainAddr];
        require(pool.mainAddr == msg.sender || pool.payerAddr == msg.sender);
        pool.updateBasic(name, email, url);
        vm.upgradeVer(Entity.ActionType.PoolUpdate);
    }

    bytes32 public CurMsgHash;
    address public RecoveredAddr;
    function TestClaimSig(address contractAddr, address tokenAddr, address user, address pool, uint amount, uint micNonce, uint cn, bytes memory signature) public{
        bytes32 message = keccak256(abi.encode(contractAddr, tokenAddr, user, pool, amount, micNonce, cn));
        CurMsgHash = prefixed(message);
        RecoveredAddr = recoverSigner(CurMsgHash, signature);
    }

    function Claim(address user, address pool, uint credit, uint amount, uint micNonce, uint cn, bytes memory signature) public{
        uint total = amount + credit;
        require(total > 0);
        require(PoolData[pool].mainAddr == msg.sender || PoolData[pool].payerAddr == msg.sender);

        Entity.UserDetails storage ud = UserData[user][pool];
        require(ud.nonce > 0);
        require(ud.epoch == cn);
        require(ud.claimedMicNonce < micNonce);

        bytes32 message = keccak256(abi.encode(this, token, user, pool, credit, amount, micNonce, cn));
        bytes32 msgHash = prefixed(message);
        require(recoverSigner(msgHash, signature) == user);

        uint tn = total.mul(1 szabo).div(setting.MBytesPerToken);
        if (tn > ud.tokenBalance){
            tn = ud.tokenBalance;
        }
        token.transfer(msg.sender, tn);
        ud.expiration = now + setting.Duration;
        ud.claim(total, tn , micNonce);

        vm.upgradeVer(Entity.ActionType.PoolClaim);
        emit PoolClaim(pool, user, total, tn, micNonce, cn);
    }

    /// builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encode("\x19Ethereum Signed Message:\n32", hash));
    }
    function recoverSigner(bytes32 message, bytes memory sig) internal pure  returns (address) {
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);
        return ecrecover(message, v, r, s);
    }
    /// signature methods.
    function splitSignature(bytes memory sig) internal pure returns (uint8 v, bytes32 r, bytes32 s) {
        require(sig.length == 65);
        assembly {
        // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
        // second 32 bytes.
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }
        return (v, r, s);
    }

    /******************
    * user action
    ******************/
    function BuyPacket(address user, uint tokenNo, address poolAddr) public{

        require(PoolData[poolAddr].mainAddr != address(0));

        token.transferFrom(msg.sender, address(this), tokenNo);

        uint newPackets = tokenNo.div(1 szabo).mul(setting.MBytesPerToken);

        Entity.UserDetails storage ch = UserData[user][poolAddr];
        if (ch.nonce == 0){
            poolsUsedByUser[user].push(poolAddr);
            usersUnderPool[poolAddr].push(user);
        }

        ch.updateChannel(tokenNo, newPackets, setting.Duration);
        vm.upgradeVer(Entity.ActionType.UserRecharge);
        emit UserRecharge(user, poolAddr, tokenNo, newPackets);
    }

    function AllMyPools(address userAddress) public view returns (address[] memory){
        return poolsUsedByUser[userAddress];
    }

    /******************
    * miner action
    ******************/
    function JoinPool(address pool, uint tokenNo,  bytes32 subAddr, bytes2 zone) public{
        require(tokenNo >= setting.MinMinerGuarantee);
        require(MinerData[subAddr].payer == address(0));

        token.transferFrom(msg.sender, address(this), tokenNo);
        Entity.Miner memory miner = Entity.Miner(minersUnderPool[pool].length, msg.sender, pool, subAddr, tokenNo, zone);

        minersUnderPool[pool].push(subAddr);
        MinerData[subAddr] = miner;
        vm.upgradeVer(Entity.ActionType.MinerJoin);
    }

    function ChangePool(address from, address to, bytes32 subAddr) public{
        Entity.Miner storage miner = MinerData[subAddr];
        require(miner.payer == msg.sender);
        require(from != to);

        delete minersUnderPool[from][miner.ID];
        miner.ID = minersUnderPool[to].length;
        miner.poolAddr = to;
        minersUnderPool[to].push(subAddr);
        vm.upgradeVer(Entity.ActionType.MinerChanged);
    }

    function RetireFromPool(bytes32 subAddr) public{
        Entity.Miner memory miner = MinerData[subAddr];
        require(miner.payer == msg.sender);

        delete minersUnderPool[miner.poolAddr][miner.ID];
        delete MinerData[subAddr];

        token.transfer(msg.sender, miner.GTN);
        vm.upgradeVer(Entity.ActionType.MinerRetire);
    }

    function AllMinersOfPool(address pool) public view returns(bytes32[] memory) {
        return minersUnderPool[pool];
    }

    //[start, end)
    function PartOfMiners(address pool, uint start, uint end) public view returns(bytes32[] memory) {
        require(end > start);
        uint length = minersUnderPool[pool].length;
        if (end > length){
            end = length;
        }
        bytes32[] memory tmp = new bytes32[](end - start);
        for (uint i = start; i < end; i++){
            tmp[i -start] = minersUnderPool[pool][i];
        }

        return tmp;
    }

    function MinerNoOfPool(address pool)public view returns (uint){
        return minersUnderPool[pool].length;
    }
}