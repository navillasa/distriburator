package main

import (
	"log"
	"os"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	word1 := []string{
		"old",
		"atomic",
		"byzantine",
		"thundering",
		"big",
		"super",
		"state",
		"shared",
		"faulty",
		"hierarchical",
		"transitive",
		"three",
		"terminating",
		"serial",
		"logical",
		"concurrent",
		"unreliable",
		"reliable",
		"persistent",
		"quantum",
		"timing",
		"automatic",
		"cryptographic",
		"sharded",
		"shared",
		"continuous",
		"self",
		"replicated",
		"global",
		"good ol'",
		"dead",
	}
	word2 := []string{
		"herd",
		"horse",
		"fault",
		"race",
		"stabilization",
		"generals",
		"data",
		"leader",
		"concurrency",
		"recovery",
		"machine",
		"snapshot",
		"version",
		"algorithm",
		"availability",
		"memory",
		"commit",
		"clock",
		"graph",
		"redundancy",
		"uniform",
		"performance",
		"coherent",
		"replication",
		"transaction",
		"object",
		"vector",
		"state",
		"conflict",
	}
	word3 := []string{
		"problem",
		"conundrum",
		"failure",
		"fallacy",
		"consensus",
		"tolerance",
		"complexity",
		"recovery",
		"law",
		"election",
		"lineage",
		"architecture",
		"tolerance",
		"theory",
		"theorem",
		"mitigation",
		"handling",
		"model",
		"deadlock",
		"synchronization",
		"control",
		"process",
		"management",
		"consistency",
		"model",
		"virtualization",
		"paradigm",
		"agreement",
		"scheme",
		"protocol",
		"cycle",
		"pitfall",
		"architecture",
		"transparency",
		"skew",
		"computing",
		"coordination",
		"pattern",
	}
	rand.Seed(time.Now().Unix())

	numTerms := 1
	var err error
	if len(os.Args) > 1 {
		numTerms, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("please use integer for number of terms you want as arg look i'm not a genius")
		}
	}
	for i:=0;i<numTerms;i++ {
		fmt.Printf("%s %s %s\n", word1[rand.Intn(len(word1))], word2[rand.Intn(len(word2))], word3[rand.Intn(len(word3))])
	}
}

