# Grin 'n Grill Backend (Go)

Welcome to the Grin 'n Grill Backend project built in Golang! This repository contains the backend infrastructure for Grin 'n Grill, a Naija-inspired culinary haven.

![grit](https://github.com/byte-man74/Grit_n-_Griddle/assets/80783021/355edf73-0c0d-4884-a1d4-4fdf29c8ab1a)


## Overview

This backend system, developed in Golang, provides a seamless and secure experience for handling online orders, payment processing, and customer management. It serves as the backbone supporting the tech-driven feast of joy at Grin 'n Grill.

## System Design and Architecture

For a detailed view of the system design and architecture, please check the [System Design Document](https://www.figma.com/file/r4ytkkknzlOgmbP2bFTHnM/Grin-n'-GRiddle?type=whiteboard&node-id=0%3A1&t=uZMKGDKBBGG9pXWR-1).

## Features

- **Order Management:** Efficient handling of online orders to streamline the restaurant's workflow.
- **Payment Processing:** Secure and reliable processing of online payments for a hassle-free transaction experience.
- **Customer Management:** A system to manage customer data, ensuring personalized and efficient service.

## Setup Instructions

1. Clone this repository: `git clone https://github.com/yourusername/grin-n-grill-backend-golang.git`
2. Install dependencies: `go mod tidy`
3. Configure environment variables.
4. Build and run the backend server: `go run main.go`

## Technologies Used

- Golang
- Your preferred database (Redis, PostgreSQL)
- Payment Gateway Integration (Opay)


## Contribution Guidelines

We welcome contributions! Feel free to open issues or submit pull requests to improve the functionality and efficiency of the Grin 'n Grill backend written in Golang.

## License

This project is licensed under the [MIT License](LICENSE.md).

Grit_n-_Griddle
├─ .git
│  ├─ COMMIT_EDITMSG
│  ├─ FETCH_HEAD
│  ├─ HEAD
│  ├─ ORIG_HEAD
│  ├─ branches
│  ├─ config
│  ├─ description
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ main
│  │     └─ remotes
│  │        └─ origin
│  │           ├─ HEAD
│  │           └─ main
│  ├─ objects
│  │  ├─ 03
│  │  │  └─ 88c9c1f7f955763e0ee5c6486bec03140fd766
│  │  ├─ 04
│  │  │  ├─ 86bb3b4673a874ddf402a75ad092ef3bead14f
│  │  │  └─ a606eb93b8856fb6dd7a7ccbc43b7a59ef8f91
│  │  ├─ 07
│  │  │  └─ 3fbaac5c93ceb5a085e3dc7e85d7a09caf0aff
│  │  ├─ 08
│  │  │  └─ 1f327738a0818048165f1b884f698251df0828
│  │  ├─ 10
│  │  │  └─ 81a3d7a99bcfa3f2d74f41d9c9bda45b4b5e40
│  │  ├─ 11
│  │  │  └─ 2ad7acb85427fae85eb2b110296d0f8886f191
│  │  ├─ 17
│  │  │  └─ 033a69fdd704d48403739250f08dd74a68e856
│  │  ├─ 1c
│  │  │  └─ a744f4b5f916201ce324d7d2709682ef984a20
│  │  ├─ 1f
│  │  │  └─ 4edd614ff77e644488bae7c36d10aaae9bc23f
│  │  ├─ 22
│  │  │  └─ da511d9c98db55ea103334f8c6ab1f0ea0be1c
│  │  ├─ 24
│  │  │  ├─ 1a8f80a29ae19c877b74b8f63248eedde1d015
│  │  │  └─ 304a51f494849c1bc14de903829ec5bfb0cffa
│  │  ├─ 27
│  │  │  └─ c3973868537a872631e2740889794c01d0e867
│  │  ├─ 2c
│  │  │  └─ 3a032ac30c9294b3ecdf2ff7732253d9eb96ac
│  │  ├─ 31
│  │  │  └─ 1931aa6924d4f13146a7f943f0f2a9b6d4308e
│  │  ├─ 39
│  │  │  └─ bc7b0dcb7494c302447c1d49ebba4f3b4954a0
│  │  ├─ 3b
│  │  │  └─ c6dc098248dda570e32211a292a6a74239ddd5
│  │  ├─ 3c
│  │  │  ├─ 7af3aa7a218eb410c99b43358cc679dd943497
│  │  │  └─ 7b8b018d72ed265c00655e7a543e533efd680c
│  │  ├─ 3d
│  │  │  └─ 9214bed52781847b3f4595d4708a3e1686d9c7
│  │  ├─ 3e
│  │  │  └─ 5562489dac14bd821a6ed0fda203af26a7f459
│  │  ├─ 41
│  │  │  └─ ba5bf13b5a2c909a324dcdeee21e2e6af7539f
│  │  ├─ 43
│  │  │  ├─ a34134eb5194080619d91d210065373f7ead25
│  │  │  └─ fcb414adce06332f28dfddd3f1d3cbd2757a9e
│  │  ├─ 44
│  │  │  └─ 219f265334dc989ff61379744a146cc4e168f8
│  │  ├─ 46
│  │  │  └─ f39fd38705d11e5174db1d0e173a5df54ad0b4
│  │  ├─ 47
│  │  │  ├─ 2fdffbdbd7f268f32b80d9b88ae958c21615d2
│  │  │  └─ 95ef8d29f473795a5941f5666f9d85b1d8a84a
│  │  ├─ 48
│  │  │  ├─ 143671505c5f1a96ecd6935a7a03ad890f222b
│  │  │  └─ 256a96f1790a95e215ac3adcace219e6015fa0
│  │  ├─ 4c
│  │  │  ├─ 15849caf3742cbf439f2ca3222e3af5b52b82c
│  │  │  └─ e5402546c4e1e6cca11bd85bead8ab4d2bd5cd
│  │  ├─ 54
│  │  │  └─ febb8e7d9d3227dd4445efe99869ba0b4d47da
│  │  ├─ 59
│  │  │  └─ 08d186c814410d5fc09820b29e3c112e4e5aca
│  │  ├─ 5b
│  │  │  └─ c6e8389591c92dee201a9fc4cfb88ebda17c25
│  │  ├─ 5e
│  │  │  └─ e3d91d1d6038c6e7022882d8fc1806916d1a4d
│  │  ├─ 5f
│  │  │  └─ e6db9574f407254425fc00ee5ca804bf3a3760
│  │  ├─ 65
│  │  │  └─ b0a91c904282ac2c71e7da8781dd7088f31267
│  │  ├─ 67
│  │  │  └─ e43ee295fb32bbc7ad858a8788f6593544b439
│  │  ├─ 69
│  │  │  └─ 1a94fa2bb47c61a5e453207fe9375a7360f258
│  │  ├─ 71
│  │  │  └─ b2b6f1fcbe2badae1b882ccf2aed7d9ef91a9a
│  │  ├─ 76
│  │  │  └─ a69fbf037dadada9c3d991ada08eedccc067c0
│  │  ├─ 79
│  │  │  └─ 249b55416af715c36a71bb157ed61c3d3f86fa
│  │  ├─ 7a
│  │  │  ├─ 26a1e951e5a1303440e6ff3796290d36053f18
│  │  │  └─ 965ea097f789665123d5b7fce69d2356fd2c9c
│  │  ├─ 7c
│  │  │  └─ d395fa3eb9bd55e710c8d37b894fe44843b3cf
│  │  ├─ 80
│  │  │  ├─ 115b717af7d2b4ac72a4598edbfd448e1cd024
│  │  │  └─ a2742376d1c717215b0c3ad99a6b5ffd141fa2
│  │  ├─ 83
│  │  │  └─ b078a139e7be47437c7ea836b0b5be6801a348
│  │  ├─ 88
│  │  │  └─ d4a9037b10bf0707f203b9424f9df0c1a78ecc
│  │  ├─ 8e
│  │  │  └─ 2ac37cfb11020e4ac41c59078d99e35b563cba
│  │  ├─ 93
│  │  │  └─ c628e091f43c20d949947a676f10a7d385a594
│  │  ├─ 95
│  │  │  └─ f5d65640898bc1245b7ad8111de092c887e25c
│  │  ├─ 9d
│  │  │  └─ 99db507d094049e3ef30b4d95aa930e9967f6d
│  │  ├─ 9e
│  │  │  └─ 55f72542e05890262fcabf8d74e25223db5e9c
│  │  ├─ a3
│  │  │  └─ 4a2fcc9e7ca72a4cff09887c0fd6457396b666
│  │  ├─ a6
│  │  │  └─ 57815c05e12f30cbc3d9e6539aac25a6152f81
│  │  ├─ a7
│  │  │  └─ eb17dc2d3b8759086011a9ba2714ee760633b6
│  │  ├─ a8
│  │  │  └─ d2f05a9fb097e14f35f49d50bfdbca64c6f8f7
│  │  ├─ ab
│  │  │  └─ 683920eaac9a4504c3a8c4c9d4a5fc7e5871e8
│  │  ├─ ad
│  │  │  └─ d3ff4558c283b66d81e445375e0e6ac05b723c
│  │  ├─ ae
│  │  │  └─ 41868afa25039defe3dc44e95f8f479680b189
│  │  ├─ b7
│  │  │  └─ 9bae4fc5b4361e60a1a5d562d9b7c3c4747248
│  │  ├─ ba
│  │  │  └─ e062236e5c524f83c885752e595dbf68a7ea09
│  │  ├─ c0
│  │  │  ├─ 34bf916d7d141a29b2f4e1e3173b12c019bd34
│  │  │  └─ efffa7a17fedb91ad6e5f0cac3ee859bd1b2c7
│  │  ├─ c7
│  │  │  ├─ 689296771e50882223373e37aef7d85625bc19
│  │  │  └─ cbc6f63c2b0b3b50191a3fdeab2d712d78f876
│  │  ├─ c9
│  │  │  └─ 03527cf8f3210b113380563e29f12b32401d00
│  │  ├─ cd
│  │  │  └─ 5063cec4cda67c2bfc3c09c454379560c1b515
│  │  ├─ ce
│  │  │  └─ eef5c2110f334e113cc042bc30e478b180ff7f
│  │  ├─ d2
│  │  │  └─ b6f59d002d3d1486147d7a2fe3348d6f6ccf15
│  │  ├─ d4
│  │  │  └─ d7142644bca53ee5dcf3df60b77cfd91b56043
│  │  ├─ d7
│  │  │  ├─ 0c00a68b7c420500594c55d5cd0d38f7824d46
│  │  │  ├─ ee9165da1318b6b12da07fc47aeca40aa333aa
│  │  │  └─ f6d8c7cece2f255bf993f4f37658bd529815e0
│  │  ├─ d9
│  │  │  └─ ecf3b414ea28e32455d5708e55c7c7631dbd96
│  │  ├─ dd
│  │  │  └─ b3c2d77642697f062903874a3d3e0382bcf6ae
│  │  ├─ e2
│  │  │  └─ cb0ee79c263c6844a124d8700aa1c3fbf4ae91
│  │  ├─ e6
│  │  │  └─ 9de29bb2d1d6434b8b29ae775ad8c2e48c5391
│  │  ├─ f3
│  │  │  ├─ 0a36d426589777c5bb9b34a6398eee36ccfd8c
│  │  │  └─ 70e0b856898c529db723026628e3e643880fec
│  │  ├─ f6
│  │  │  └─ c73294e1b0a5a88061cf69153d2b67376998ea
│  │  ├─ f7
│  │  │  └─ 24779c3bbc9b4eca574d33d41b190c29fe8311
│  │  ├─ fa
│  │  │  └─ b9579c884b264856116c24e4f154ebf05110ae
│  │  ├─ fe
│  │  │  └─ aa7fdb339edfb1b1cbeb629c7c46a9a6834729
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-552371bc5f5ea69c75e1bd30f1f733f9c13e52c8.idx
│  │     └─ pack-552371bc5f5ea69c75e1bd30f1f733f9c13e52c8.pack
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ main
│     ├─ remotes
│     │  └─ origin
│     │     ├─ HEAD
│     │     └─ main
│     └─ tags
├─ .gitignore
├─ LICENSE
├─ README.md
├─ backend
│  ├─ controllers
│  │  ├─ authentication_controller.go
│  │  └─ whatisthis.md
│  ├─ database
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  ├─ middleware
│  │  └─ whatisthis.md
│  ├─ models
│  │  └─ userModel.go
│  ├─ routes
│  └─ utils
│     └─ authentication_utils
│        └─ user_utils.go
└─ reference.md


Grit_n-_Griddle
├─ .git
│  ├─ COMMIT_EDITMSG
│  ├─ FETCH_HEAD
│  ├─ HEAD
│  ├─ ORIG_HEAD
│  ├─ branches
│  ├─ config
│  ├─ description
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ main
│  │     └─ remotes
│  │        └─ origin
│  │           ├─ HEAD
│  │           └─ main
│  ├─ objects
│  │  ├─ 01
│  │  │  └─ 85fdb0e31222cc96b27b4e189f7550c2ad8f56
│  │  ├─ 03
│  │  │  └─ 88c9c1f7f955763e0ee5c6486bec03140fd766
│  │  ├─ 04
│  │  │  ├─ 86bb3b4673a874ddf402a75ad092ef3bead14f
│  │  │  └─ a606eb93b8856fb6dd7a7ccbc43b7a59ef8f91
│  │  ├─ 07
│  │  │  └─ 3fbaac5c93ceb5a085e3dc7e85d7a09caf0aff
│  │  ├─ 08
│  │  │  └─ 1f327738a0818048165f1b884f698251df0828
│  │  ├─ 0a
│  │  │  ├─ 729dadd69e9b4a7885ef1fc86f385f61c63e37
│  │  │  └─ 9d8add26f39a41bf5d6bc856ae485366a50511
│  │  ├─ 10
│  │  │  └─ 81a3d7a99bcfa3f2d74f41d9c9bda45b4b5e40
│  │  ├─ 11
│  │  │  └─ 2ad7acb85427fae85eb2b110296d0f8886f191
│  │  ├─ 14
│  │  │  └─ f8de485e1e6a6ae41e5204953000b6587ab533
│  │  ├─ 17
│  │  │  └─ 033a69fdd704d48403739250f08dd74a68e856
│  │  ├─ 1c
│  │  │  └─ a744f4b5f916201ce324d7d2709682ef984a20
│  │  ├─ 1f
│  │  │  └─ 4edd614ff77e644488bae7c36d10aaae9bc23f
│  │  ├─ 20
│  │  │  └─ 3db845ef02f3807a124f214642a148a4843083
│  │  ├─ 22
│  │  │  └─ da511d9c98db55ea103334f8c6ab1f0ea0be1c
│  │  ├─ 24
│  │  │  ├─ 1a8f80a29ae19c877b74b8f63248eedde1d015
│  │  │  └─ 304a51f494849c1bc14de903829ec5bfb0cffa
│  │  ├─ 27
│  │  │  └─ c3973868537a872631e2740889794c01d0e867
│  │  ├─ 2c
│  │  │  └─ 3a032ac30c9294b3ecdf2ff7732253d9eb96ac
│  │  ├─ 31
│  │  │  └─ 1931aa6924d4f13146a7f943f0f2a9b6d4308e
│  │  ├─ 32
│  │  │  └─ 6829a3ca9ddcea9256abed5c9652867ed2655d
│  │  ├─ 35
│  │  │  └─ 83daae61997adc5f3df726bcc460a729926090
│  │  ├─ 39
│  │  │  ├─ 09f851032e30acb943b3f311e7d3e5c5f21847
│  │  │  └─ bc7b0dcb7494c302447c1d49ebba4f3b4954a0
│  │  ├─ 3b
│  │  │  └─ c6dc098248dda570e32211a292a6a74239ddd5
│  │  ├─ 3c
│  │  │  ├─ 7af3aa7a218eb410c99b43358cc679dd943497
│  │  │  └─ 7b8b018d72ed265c00655e7a543e533efd680c
│  │  ├─ 3d
│  │  │  └─ 9214bed52781847b3f4595d4708a3e1686d9c7
│  │  ├─ 3e
│  │  │  └─ 5562489dac14bd821a6ed0fda203af26a7f459
│  │  ├─ 41
│  │  │  └─ ba5bf13b5a2c909a324dcdeee21e2e6af7539f
│  │  ├─ 43
│  │  │  ├─ a34134eb5194080619d91d210065373f7ead25
│  │  │  └─ fcb414adce06332f28dfddd3f1d3cbd2757a9e
│  │  ├─ 44
│  │  │  └─ 219f265334dc989ff61379744a146cc4e168f8
│  │  ├─ 46
│  │  │  └─ f39fd38705d11e5174db1d0e173a5df54ad0b4
│  │  ├─ 47
│  │  │  ├─ 2fdffbdbd7f268f32b80d9b88ae958c21615d2
│  │  │  └─ 95ef8d29f473795a5941f5666f9d85b1d8a84a
│  │  ├─ 48
│  │  │  ├─ 143671505c5f1a96ecd6935a7a03ad890f222b
│  │  │  └─ 256a96f1790a95e215ac3adcace219e6015fa0
│  │  ├─ 4c
│  │  │  ├─ 15849caf3742cbf439f2ca3222e3af5b52b82c
│  │  │  └─ e5402546c4e1e6cca11bd85bead8ab4d2bd5cd
│  │  ├─ 54
│  │  │  └─ febb8e7d9d3227dd4445efe99869ba0b4d47da
│  │  ├─ 59
│  │  │  ├─ 08d186c814410d5fc09820b29e3c112e4e5aca
│  │  │  └─ 367a99fcca0a90f929350de5746414b7b663c2
│  │  ├─ 5b
│  │  │  └─ c6e8389591c92dee201a9fc4cfb88ebda17c25
│  │  ├─ 5c
│  │  │  ├─ 471aba4d966a67bd3da83562c3acf16b52d933
│  │  │  └─ fa5cf3827d3ecebc85e0fd1b6d955007282af4
│  │  ├─ 5d
│  │  │  └─ 65d491ba976fb415650c5b999e8eeb84e9c733
│  │  ├─ 5e
│  │  │  └─ e3d91d1d6038c6e7022882d8fc1806916d1a4d
│  │  ├─ 5f
│  │  │  └─ e6db9574f407254425fc00ee5ca804bf3a3760
│  │  ├─ 65
│  │  │  ├─ 1c02066a8a91b1aed8a540f80151cc66fa0ee1
│  │  │  └─ b0a91c904282ac2c71e7da8781dd7088f31267
│  │  ├─ 67
│  │  │  └─ e43ee295fb32bbc7ad858a8788f6593544b439
│  │  ├─ 69
│  │  │  └─ 1a94fa2bb47c61a5e453207fe9375a7360f258
│  │  ├─ 71
│  │  │  └─ b2b6f1fcbe2badae1b882ccf2aed7d9ef91a9a
│  │  ├─ 74
│  │  │  └─ 027fc2fe1a5ee2a08037def5c4b5e16bf79bfa
│  │  ├─ 76
│  │  │  └─ a69fbf037dadada9c3d991ada08eedccc067c0
│  │  ├─ 78
│  │  │  └─ 7581411f867ca85cb0888f05411f9209fe8cc9
│  │  ├─ 79
│  │  │  └─ 249b55416af715c36a71bb157ed61c3d3f86fa
│  │  ├─ 7a
│  │  │  ├─ 26a1e951e5a1303440e6ff3796290d36053f18
│  │  │  └─ 965ea097f789665123d5b7fce69d2356fd2c9c
│  │  ├─ 7c
│  │  │  └─ d395fa3eb9bd55e710c8d37b894fe44843b3cf
│  │  ├─ 7e
│  │  │  └─ 9aee2770e37c357dcc545d2d728d4d2549be91
│  │  ├─ 80
│  │  │  ├─ 115b717af7d2b4ac72a4598edbfd448e1cd024
│  │  │  └─ a2742376d1c717215b0c3ad99a6b5ffd141fa2
│  │  ├─ 83
│  │  │  └─ b078a139e7be47437c7ea836b0b5be6801a348
│  │  ├─ 88
│  │  │  └─ d4a9037b10bf0707f203b9424f9df0c1a78ecc
│  │  ├─ 8d
│  │  │  └─ 410eddc9324e8a0eac378f6e0f899e3646da63
│  │  ├─ 8e
│  │  │  └─ 2ac37cfb11020e4ac41c59078d99e35b563cba
│  │  ├─ 90
│  │  │  └─ 10a9a0431eb4bad7626139acfbc095b297199c
│  │  ├─ 93
│  │  │  └─ c628e091f43c20d949947a676f10a7d385a594
│  │  ├─ 95
│  │  │  └─ f5d65640898bc1245b7ad8111de092c887e25c
│  │  ├─ 96
│  │  │  └─ f67b08b790054b405a1ee940db1e94698ced6a
│  │  ├─ 99
│  │  │  └─ 18021106bab5e0d5c79a7961ce5bc0790d0a93
│  │  ├─ 9d
│  │  │  └─ 99db507d094049e3ef30b4d95aa930e9967f6d
│  │  ├─ 9e
│  │  │  └─ 55f72542e05890262fcabf8d74e25223db5e9c
│  │  ├─ a0
│  │  │  └─ 57acf385f87b5bbaf15766a1c426935d119351
│  │  ├─ a2
│  │  │  └─ f3e5776bb550ffefd03e5b86afdf8f9f18ec3c
│  │  ├─ a3
│  │  │  └─ 4a2fcc9e7ca72a4cff09887c0fd6457396b666
│  │  ├─ a4
│  │  │  └─ 65d0542f397c02f9ec3fc5e36c32c5cc3ad508
│  │  ├─ a6
│  │  │  └─ 57815c05e12f30cbc3d9e6539aac25a6152f81
│  │  ├─ a7
│  │  │  └─ eb17dc2d3b8759086011a9ba2714ee760633b6
│  │  ├─ a8
│  │  │  └─ d2f05a9fb097e14f35f49d50bfdbca64c6f8f7
│  │  ├─ ab
│  │  │  └─ 683920eaac9a4504c3a8c4c9d4a5fc7e5871e8
│  │  ├─ ad
│  │  │  └─ d3ff4558c283b66d81e445375e0e6ac05b723c
│  │  ├─ ae
│  │  │  └─ 41868afa25039defe3dc44e95f8f479680b189
│  │  ├─ b7
│  │  │  └─ 9bae4fc5b4361e60a1a5d562d9b7c3c4747248
│  │  ├─ ba
│  │  │  └─ e062236e5c524f83c885752e595dbf68a7ea09
│  │  ├─ c0
│  │  │  ├─ 34bf916d7d141a29b2f4e1e3173b12c019bd34
│  │  │  └─ efffa7a17fedb91ad6e5f0cac3ee859bd1b2c7
│  │  ├─ c1
│  │  │  └─ ef97d9d4296eec1b55b275e825d13e8c9750b0
│  │  ├─ c7
│  │  │  ├─ 689296771e50882223373e37aef7d85625bc19
│  │  │  └─ cbc6f63c2b0b3b50191a3fdeab2d712d78f876
│  │  ├─ c9
│  │  │  └─ 03527cf8f3210b113380563e29f12b32401d00
│  │  ├─ cd
│  │  │  └─ 5063cec4cda67c2bfc3c09c454379560c1b515
│  │  ├─ ce
│  │  │  └─ eef5c2110f334e113cc042bc30e478b180ff7f
│  │  ├─ d2
│  │  │  └─ b6f59d002d3d1486147d7a2fe3348d6f6ccf15
│  │  ├─ d4
│  │  │  └─ d7142644bca53ee5dcf3df60b77cfd91b56043
│  │  ├─ d5
│  │  │  └─ f5d551c3893585faffae1bf32dc630f626d982
│  │  ├─ d6
│  │  │  └─ fa8e37649ea44d93f4d5edc16c9043a8825483
│  │  ├─ d7
│  │  │  ├─ 0c00a68b7c420500594c55d5cd0d38f7824d46
│  │  │  ├─ ee9165da1318b6b12da07fc47aeca40aa333aa
│  │  │  └─ f6d8c7cece2f255bf993f4f37658bd529815e0
│  │  ├─ d9
│  │  │  └─ ecf3b414ea28e32455d5708e55c7c7631dbd96
│  │  ├─ dd
│  │  │  └─ b3c2d77642697f062903874a3d3e0382bcf6ae
│  │  ├─ e2
│  │  │  └─ cb0ee79c263c6844a124d8700aa1c3fbf4ae91
│  │  ├─ e6
│  │  │  ├─ 629815c9f2d0d8d6eff8a6ba96ae99415396f1
│  │  │  └─ 9de29bb2d1d6434b8b29ae775ad8c2e48c5391
│  │  ├─ ef
│  │  │  ├─ 158f47a47a4d0ee12040a81d7dff1ad46636f9
│  │  │  └─ 648ecb5f074aae4852868075778147cfb49952
│  │  ├─ f3
│  │  │  ├─ 0a36d426589777c5bb9b34a6398eee36ccfd8c
│  │  │  └─ 70e0b856898c529db723026628e3e643880fec
│  │  ├─ f6
│  │  │  ├─ 0121571a3cf7ee59db57c493406df4ce9e73b2
│  │  │  ├─ 6d3e72e2eb4b16804804317138beea77e79108
│  │  │  ├─ 8118454cf69ceeba11814763e7173dbfa40b51
│  │  │  └─ c73294e1b0a5a88061cf69153d2b67376998ea
│  │  ├─ f7
│  │  │  └─ 24779c3bbc9b4eca574d33d41b190c29fe8311
│  │  ├─ f9
│  │  │  └─ 96eb48ffee3653734a1af1ac71e075c97014b2
│  │  ├─ fa
│  │  │  └─ b9579c884b264856116c24e4f154ebf05110ae
│  │  ├─ fe
│  │  │  └─ aa7fdb339edfb1b1cbeb629c7c46a9a6834729
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-552371bc5f5ea69c75e1bd30f1f733f9c13e52c8.idx
│  │     └─ pack-552371bc5f5ea69c75e1bd30f1f733f9c13e52c8.pack
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ main
│     ├─ remotes
│     │  └─ origin
│     │     ├─ HEAD
│     │     └─ main
│     └─ tags
├─ .gitignore
├─ LICENSE
├─ README.md
├─ backend
│  ├─ controllers
│  │  └─ whatisthis.md
│  ├─ database
│  ├─ go.mod
│  ├─ go.sum
│  ├─ initializers
│  │  ├─ database.go
│  │  └─ load_env.go
│  ├─ main.go
│  ├─ middleware
│  │  └─ whatisthis.md
│  ├─ migrate
│  │  └─ migrate.go
│  ├─ models
│  │  └─ userModel.go
│  ├─ routes
│  └─ utils
│     └─ authentication_utils
│        └─ user_utils.go
└─ reference.md

```