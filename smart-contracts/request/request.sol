pragma solidity ^0.4.0;
contract Request {

    struct Contribution {
    uint id;
    bool passedInspection;
    }

    address public requester;

    Contribution[] contribution;

    function Request() {
        requester = msg.sender;
    }
}