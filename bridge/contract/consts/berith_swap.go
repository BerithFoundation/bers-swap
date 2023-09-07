package consts

const BerithSwapABI = `
[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint64",
          "name": "depositNonce",
          "type": "uint64"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "receipient",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "Deposit",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "receipientAddress",
          "type": "address"
        }
      ],
      "name": "deposit",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "depositNonce",
      "outputs": [
        {
          "internalType": "uint64",
          "name": "",
          "type": "uint64"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "transferFunds",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]
`

const BerithSwapBin = "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6107998061010d6000396000f3fe6080604052600436106100555760003560e01c80633c68eb811461005a578063715018a6146100715780638da5cb5b14610088578063de35f5cb146100b3578063f2fde38b146100de578063f340fa0114610107575b600080fd5b34801561006657600080fd5b5061006f610123565b005b34801561007d57600080fd5b50610086610181565b005b34801561009457600080fd5b5061009d610195565b6040516100aa9190610566565b60405180910390f35b3480156100bf57600080fd5b506100c86101be565b6040516100d591906105c1565b60405180910390f35b3480156100ea57600080fd5b50610105600480360381019061010091906104ca565b6101d8565b005b610121600480360381019061011c91906104ca565b61025c565b005b61012b61036b565b6000610135610195565b90508073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561017d573d6000803e3d6000fd5b5050565b61018961036b565b61019360006103e9565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600060149054906101000a900467ffffffffffffffff1681565b6101e061036b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610250576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024790610581565b60405180910390fd5b610259816103e9565b50565b6000806102676104ad565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156102ad576102a66104ad565b91506102b1565b8291505b6000601481819054906101000a900467ffffffffffffffff16809291906102d790610674565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508073ffffffffffffffffffffffffffffffffffffffff167fd161c4b6e400003b6ab42af5a25de2cdb081619bcf8d705f4bd8c50aa5e2543e600060149054906101000a900467ffffffffffffffff16843460405161035e939291906105dc565b60405180910390a2505050565b6103736104ad565b73ffffffffffffffffffffffffffffffffffffffff16610391610195565b73ffffffffffffffffffffffffffffffffffffffff16146103e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103de906105a1565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000813590506104c48161074c565b92915050565b6000602082840312156104dc57600080fd5b60006104ea848285016104b5565b91505092915050565b6104fc81610624565b82525050565b600061050f602683610613565b915061051a826106d4565b604082019050919050565b6000610532602083610613565b915061053d82610723565b602082019050919050565b61055181610656565b82525050565b61056081610660565b82525050565b600060208201905061057b60008301846104f3565b92915050565b6000602082019050818103600083015261059a81610502565b9050919050565b600060208201905081810360008301526105ba81610525565b9050919050565b60006020820190506105d66000830184610557565b92915050565b60006060820190506105f16000830186610557565b6105fe60208301856104f3565b61060b6040830184610548565b949350505050565b600082825260208201905092915050565b600061062f82610636565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061067f82610660565b915067ffffffffffffffff82141561069a576106996106a5565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b61075581610624565b811461076057600080fd5b5056fea2646970667358221220284e3a6e563e01caf7c01db1c6a48a4fef08ba585fd09366ea4bfd1f828f700164736f6c63430008010033"
