package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type IDtoFIslice map[int][]int

// seq_name sample_id sample_type organism donor_id sex age_days
//  brain_hemisphere brain_region brain_subregion facs_date facs_container
//  sample_name facs_sort_criteria rna_amplification_set library_prep_set
//  library_prep_avg_size_bp seq_tube seq_batch total_reads percent_exon_reads
//  percent_intron_reads percent_intergenic_reads percent_rrna_reads
//  percent_unique_reads percent_synth_reads percent_ecoli_reads
//  percent_aligned_reads_total complexity_cg genes_detected_cpm_criterion
//  genes_detected_fpkm_criterion]
//  [324 SM-GE8ZP R8S4-180117 1304148 51.37070647 37.48830223 11.14099131 0.053904925 83.00215926 0.015656965 0.003363882 91.65846208 0.397504042 6102 3256]
// 	seq_name	SM-GE8ZP_S159_E1-50
//	sample_id	658376676
//	sample_type	Nuclei
//	organism	Human
//	donor_id	H200.1023
//	sex		F
//	age_days	15695
//	brain_hemisphere	L
//	brain_region	VISp
//	brain_subregion	L6b
//	facs_date	12/6/2017
//	facs_container	F1S4_171206_389
//	sample_name		F1S4_171206_389_G01
//	facs_sort_criteria	NeuN-positive
//	rna_amplification_set	A8S4_171219_07
//	library_prep_set	L8S4_180109_01
//	library_prep_avg_size_bp	324
//	seq_tube	SM-GE8ZP R8S4-180117
func ParseBarcodeFIfile(file string) (IDtoFIslice, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	idToFIslice := make(IDtoFIslice)

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'
	var headerLn []string
	for lineNum := 1; ; lineNum++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lineNum)
		fmt.Println(record)
		fmt.Println(fmt.Sprintf("%T", record))

		// The header line
		if lineNum == 1 {
			expectedHeader := []string{"barcode_id", "FI"}
			if !StringSliceEqual(record, expectedHeader) {
				return nil, fmt.Errorf("Header = %v is not what's expected (=%v)", record, expectedHeader)
			}
		} else {
			if len(record) != 2 {
				return nil, fmt.Errorf("Split line does not have 2 expected entries: %v", record)
			}
			// covert the two entries to ints
			id, err := strconv.Atoi(record[0])
			if err != nil {
				return nil, err
			}
			FI, err := strconv.Atoi(record[1])
			if err != nil {
				return nil, err
			}
			idToFIslice[id] = append(idToFIslice[id], FI)
		}
	}
	fmt.Println(headerLn)

	return idToFIslice, nil
}
