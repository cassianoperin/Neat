package main

import "fmt"

type NeuronGene struct {
	Id int
  Type string
}

var (
  population []NeuronGene
	v1 = NeuronGene{1, "Input"}  // has type Vertex
	v2 = NeuronGene{Id: 1}  // Y:0 is implicit
	v3 = NeuronGene{}      // X:0 and Y:0
	p  = &NeuronGene{2, "Output"} // has type *NeuronGene
)

func main() {
	fmt.Println(v1, p, v2, v3)

  population = append (population, NeuronGene{10, "Hidden"} )
  population = append (population, NeuronGene{20, "Bias"} )

  fmt.Printf("%d\t%s\n", population[0].Id,  population[0].Type )
  fmt.Printf("%d\t%s", population[1].Id,  population[1].Type )

}
