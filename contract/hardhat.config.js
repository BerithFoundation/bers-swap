require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-ethers");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.1",
  paths: {
    sources: "./contracts",
    artifacts: "../bridge/contract/consts/compiled",
  },
  networks: {
    berithTestnet: {
      url: "http://testnet.berith.co:8545",
      accounts: [
        "56a7a32def4a15c0187e8a076baf4c76db60bbf9ad59b8396d0b2174f9d5533e",
      ],
    },
    klaytnTestnet: {
      url: "https://api.baobab.klaytn.net:8651",
      accounts: [
        "52069ec4b9c89b5c16eff569c6f9f5d3825e07842bc40a3a833c1fab157e0646",
      ],
    },
  },
};
