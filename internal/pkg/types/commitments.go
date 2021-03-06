package types

// Commitments is a struct containing the replica and data commitments produced
// when sealing a sector.
type Commitments struct {
	CommD     CommD
	CommR     CommR
	CommRStar CommRStar
}

// PoStChallengeSeedBytesLen is the number of bytes in the Proof of SpaceTime challenge seed.
const PoStChallengeSeedBytesLen uint = 32

// CommitmentBytesLen is the number of bytes in a CommR, CommD, CommP, and CommRStar.
const CommitmentBytesLen uint = 32

// PoStChallengeSeed is an input to the proof-of-spacetime generation and verification methods.
type PoStChallengeSeed [PoStChallengeSeedBytesLen]byte

// CommR is the merkle root of the replicated data. It is an output of the
// sector sealing (PoRep) process.
type CommR [CommitmentBytesLen]byte

// CommD is the merkle root of the original user data. It is an output of the
// sector sealing (PoRep) process.
type CommD [CommitmentBytesLen]byte

// CommP is the merkle root of a piece of data included within the original user data. It is
// generated by the client, and the miner must generated a piece inclusion proof from CommP
// to CommD.
type CommP [CommitmentBytesLen]byte

// CommRStar is a hash of intermediate layers. It is an output of the sector
// sealing (PoRep) process.
type CommRStar [CommitmentBytesLen]byte
