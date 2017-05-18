pragma solidity ^0.4.11;


contract mortal {
  /* Define variable owner of the type address*/
  address owner;

  uint public created;

  uint public  modified;
  /* this function is executed at initialization and sets the owner of the contract */
  function mortal() {owner = msg.sender;}

  /* Function to recover the funds on the contract */
  function kill() {if (msg.sender == owner) selfdestruct(owner);}

  function setCreated(){
    created = block.timestamp;
  }

  function setModified(){
    modified = block.timestamp;
  }
}


contract Request is mortal {

  int64 public userId;

  address public inspectorAddress;

  uint16[] public contributionCodes;

  string public remark;


  RBAC rbac;

  struct Lack {
  uint16 contributionCode;
  string controlCategoryId;
  string pointGroupId;
  string controlPointId;
  int64 lackId;
  }

  uint public numLacks;

  mapping (uint => Lack) public lacks;

  function Request(int64 _userId, uint16[] _contributionCodes, string _remark, address rbacAddress) public {
    rbac = RBAC(rbacAddress);
    //    m.sendToken(receiver, amount);
    userId = _userId;
    contributionCodes = _contributionCodes;
    remark = _remark;
    setCreated();
  }

  function setInspectorId(address _inspectorAddress){
    //    require( rbac.isAdmin(msg.sender) );
    inspectorAddress = _inspectorAddress;
    setModified();
  }
  //
  //  function isAdmin(address inspectorAddress) constant returns (bool) {
  //    return rbac.isAdmin(inspectorAddress);
  //  }


  function addLack(uint16 _contributionCode, string _controlCategoryId, string _pointGroupId, string controlPointId, int64 lackId) {
    require(msg.sender == inspectorAddress);
    uint lacksIndex = numLacks++;
    lacks[lacksIndex] = Lack(_contributionCode, _controlCategoryId, _pointGroupId, controlPointId, lackId);
    setModified();
  }

}


/// @title Role Based Access Control
// This is basically an interface for solidity. So solidity can work out how to call the functions
contract RBAC {
  mapping (address => bool) admins;

  mapping (address => bool) inspectors;

  mapping (address => bool) farmers;

  function isFarmer(address farmerAddress) constant returns (bool) {
    return farmers[farmerAddress];
  }

  function isAdmin(address adminAddress) constant returns (bool) {
    return admins[adminAddress];
  }

  function isInspector(address inspectorAddress) constant returns (bool) {
    return inspectors[inspectorAddress];
  }
}
