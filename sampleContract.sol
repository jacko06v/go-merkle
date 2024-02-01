pragma solidity ^0.8.0;

contract HashObjectJson {
    address ethereumAddress = 0x82714c95388cFd6f366A4aBA3c4add588BBC36fC;
    uint256 number = 400000000000000000;

    function hashObjectJson() public view returns (bytes32) {
        bytes memory addressBytes = abi.encodePacked(ethereumAddress);
        string memory hexString = bytesToHexString(addressBytes);
        uint256 convertedNumber = parseInt(hexString, 16);
        bytes memory packed = abi.encodePacked(convertedNumber, number);
        return keccak256(packed);
    }

    function bytesToHexString(
        bytes memory data
    ) internal pure returns (string memory) {
        bytes memory hexString = new bytes(2 * data.length);
        for (uint256 i = 0; i < data.length; i++) {
            uint256 value = uint256(uint8(data[i]));
            bytes memory alphabet = "0123456789abcdef";
            hexString[2 * i] = alphabet[value / 16];
            hexString[2 * i + 1] = alphabet[value % 16];
        }
        return string(hexString);
    }

    function parseInt(
        string memory _value,
        uint256 _base
    ) internal pure returns (uint256) {
        bytes memory valueBytes = bytes(_value);
        uint256 result = 0;
        for (uint256 i = 0; i < valueBytes.length; i++) {
            uint256 digit = uint256(uint8(valueBytes[i]));
            if (digit >= 48 && digit <= 57) {
                digit -= 48;
            } else if (digit >= 65 && digit <= 90) {
                digit -= 55;
            } else if (digit >= 97 && digit <= 122) {
                digit -= 87;
            }
            require(digit < _base, "Invalid digit");
            result = result * _base + digit;
        }
        return result;
    }
}
