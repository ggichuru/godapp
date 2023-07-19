// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract Quiz {
    string public question;
    bytes32 internal answer;
    mapping(address => bool) internal leaderboard;

    constructor(string memory qn, bytes32 ans) {
        question = qn;
        answer = ans;
    }

    function sendAnswer(bytes32 ans) public returns (bool) {
        // call update leaderboard
        return updateLeaderboard(ans == answer);
    }

    function updateLeaderboard(bool ok) internal returns (bool) {
        // add sender to leaderboard
        leaderboard[msg.sender] = ok;
        return true;
    }

    function checkBoard() public view returns (bool) {
        // check if sender is on leaderboard
        return leaderboard[msg.sender];
    }
}
