// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

interface ITimelock {
    function executeTransaction(address target, uint value, string memory signature, bytes memory data, uint eta) external returns (bytes memory);
    function queuedTransactions(bytes32) external view returns (bool);
}

library TimelockHelper {

    bytes constant EXEC_DATA_M2_UPGRADE = hex"6A76120200000000000000000000000040A2ACCBD92BCA938B02010E17A5B8929B49130D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007A000000000000000000000000000000000000000000000000000000000000006248D80FF0A000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005D3008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000039053D51B77DC0D36036FC1FCC8CB819DF8EF37A0000000000000000000000001784BE6401339FC0FEDF7E9379409F5C1BFE9DDA008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000D92145C07F8ED1D392C1B88017934E301CC1C3CD000000000000000000000000F3234220163A757EDF1E11A8A085638D9B236614008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000858646372CC42E1A627FCE94AA7A7033E7CF075A00000000000000000000000070F44C13944D49A236E3CD7A94F48F5DAB6C619B008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000007FE7E9CC0F274D2435AD5D56D5FA73E47F6A23D80000000000000000000000004BB6731B02314D40ABBFFBC4540F508874014226008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000091E677B07F7AF907EC9A428AAFA9FC14A0D3A338000000000000000000000000E4297E3DADBC7D99E26A2954820F514CB50C5762005A2A4F2F3C18F09179B6703E63D9EDD165909073000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659CFE60000000000000000000000008BA40DA60F0827D027F029ACEE62609F0527A2550039053D51B77DC0D36036FC1FCC8CB819DF8EF37A00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024635BBD10000000000000000000000000000000000000000000000000000000000000C4E00091E677B07F7AF907EC9A428AAFA9FC14A0D3A33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024C1DE3AEF000000000000000000000000343907185B71ADF0EBA9567538314396AA9854420091E677B07F7AF907EC9A428AAFA9FC14A0D3A33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024463DB0380000000000000000000000000000000000000000000000000000000065F1B0570039053D51B77DC0D36036FC1FCC8CB819DF8EF37A00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024FABC1CBC00000000000000000000000000000000000000000000000000000000000000000091E677B07F7AF907EC9A428AAFA9FC14A0D3A33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024FABC1CBC000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000A6DB1A8C5A981D1536266D2A393C5F8DDB210EAF00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000";
    uint constant ETA_M2_UPGRADE = 1712559600;

    bytes constant EXEC_DATA_UNPAUSE_DEPOSITS = hex"6a761202000000000000000000000000858646372cc42e1a627fce94aa7a7033e7cf075a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000024fabc1cbc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000a6db1a8c5a981d1536266d2a393c5f8ddb210eaf00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000";
    uint constant ETA_UNPAUSE_DEPOSITS = 1713250800;
}