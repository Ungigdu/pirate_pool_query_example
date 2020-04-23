pragma solidity >0.5.10;

import "./safemath.sol";

library Entity {

    using SafeMath for uint256;
    enum ActionType{PoolReg, PoolUpdate, PoolClaim, PoolExit, UserRecharge, UserRefund, SysSetting, MinerJoin, MinerChanged, MinerRetire}

    struct VersionManager {
        uint versionOfPoolData;
        uint versionOfMinerData;
        uint versionOfUserData;
        uint versionOfSysSetting;
    }

    struct SystemSetting{
        uint MBytesPerToken;
        uint Duration;
        uint MinPoolGuarantee;
        uint MinMinerGuarantee;
    }

    struct UserDetails {
        uint nonce;
        uint tokenBalance;
        uint remindPacket;
        uint expiration;
        uint epoch;
        uint claimedAmount;
        uint claimedMicNonce;
    }

    struct MinerPool{
        uint ID;
        address mainAddr;
        address payerAddr;
        uint GTN;
        bytes   shortName;
        bytes   email;
        bytes   url;
    }

    struct Miner{
        uint ID;
        address payer;
        address poolAddr;
        bytes32 subAddr;
        uint GTN;
        bytes2 zone;
    }

    function upgradeVer(VersionManager storage self, ActionType at) internal{

        if (at == ActionType.PoolReg || at == ActionType.PoolUpdate || at == ActionType.PoolExit){

            self.versionOfPoolData ++;

        }else if (at == ActionType.PoolClaim || at == ActionType.UserRefund
        || at == ActionType.UserRecharge){

            self.versionOfUserData ++;

        }else if (at == ActionType.SysSetting){

            self.versionOfSysSetting ++;

        }else if (at == ActionType.MinerJoin || at == ActionType.MinerChanged
        || at == ActionType.MinerRetire){

            self.versionOfMinerData ++;

        }
    }

    function updateChannel(UserDetails storage self, uint tn, uint packs, uint d) internal {
        self.nonce += 1;
        self.tokenBalance = self.tokenBalance.add(tn);
        self.remindPacket = self.remindPacket.add(packs);
        self.expiration = now + d;
    }

    function claim(UserDetails storage self, uint amount, uint tn, uint micNonce) internal{
        self.nonce += 1;
        self.remindPacket = self.remindPacket.sub(amount);
        self.tokenBalance = self.tokenBalance.sub(tn);
        self.epoch += 1;
        self.claimedAmount = amount;
        self.claimedMicNonce = micNonce;
    }

    function updateBasic(MinerPool storage self, bytes memory name,
        bytes memory email, bytes memory url) internal {

        if (name.length != 0){
            self.shortName = name;
        }

        if (email.length != 0){
            self.email = email;
        }

        if (url.length != 0){
            self.url = url;
        }
    }
}
