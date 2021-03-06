package goscales

/*Repository can return golang byte to be mapped.*/
func Repository() []byte {
	var scaleByte = []byte(`
{
  "escala_natural":{
                  "escala_mayor_natural":       [2, 2, 1, 2, 2, 2, 1],
                  "escala_dorica":              [2, 1, 2, 2, 2, 1, 2],
                  "escala_frigia":              [1, 2, 2, 2, 1, 2, 2],
                  "escala_lidia":               [2, 2, 2, 1, 2, 2, 1],
                  "escala_mixolidia":           [2, 2, 1, 2, 2, 1, 2],
                  "escala_menor_natural":       [2, 1, 2, 2, 1, 2, 2],
                  "escala_locria":              [1, 2, 2, 1, 2, 2, 2]},


  "escala_melódica":{
                    "escala_menor_melodica":    [2, 1, 2, 2, 2, 2, 1],
                    "escala_java":              [1, 2, 2, 2, 2, 1, 2],
                    "escala_lidia_#5":          [2, 2, 2, 2, 1, 2, 1],
                    "escala_lidia_b7":          [2, 2, 2, 1, 2, 1, 2],
                    "escala_hindustan":         [2, 2, 1, 2, 1, 2, 2],
                    "escala_locria_natural_2":  [2, 1, 2, 1, 2, 2, 2],
                    "escala_alterada":          [1, 2, 1, 2, 2, 2, 2]},


  "escala_armónica_menor":{
                    "escala_menor_armonica":      [2, 1, 2, 2, 1, 3, 1],
                    "escala_locria_natural_6":    [1, 2, 2, 1, 3, 1, 2],
                    "escala_jonica_#5":           [2, 2, 1, 3, 1, 2, 1],
                    "escala_menor_romana":        [2, 1, 3, 1, 2, 1, 2],
                    "escala_frigia_mayor":        [1, 3, 1, 2, 1, 2, 2],
                    "escala_lidia_#2":            [3, 1, 2, 1, 2, 2, 1],
                    "escala_superlocria_bb7":     [1, 2, 1, 2, 2, 1, 3]},

  "escala_mixta_o_hindú":{
                    "escala_armonica_mayor":          [2, 2, 1, 2, 1, 3, 1],
                    "escala_locria_natural_2_6":      [2, 1, 2, 1, 3, 1, 2],
                    "escala_frigia_mayor_#9":         [1, 2, 1, 3, 1, 2, 2],
                    "escala_lidia_menor":             [2, 1, 3, 1, 2, 2, 1],
                    "escala_mixolidia_b2":            [1, 3, 1, 2, 2, 1, 2],
                    "escala_lidia_#2_#5":             [3, 1, 2, 2, 1, 2, 1],
                    "escala_locria_bb7":              [1, 2, 2, 1, 2, 1, 3]},

  "escala_turca":{
                    "escala_turca_primer_grado":   [1, 2, 2, 2, 2, 2, 1],
                    "escala_turca_segundo_grado":  [2, 2, 2, 2, 2, 1, 1],
                    "escala_turca_tercer_grado":   [2, 2, 2, 2, 1, 1, 2],
                    "escala_turca_cuarto_grado":   [2, 2, 2, 1, 1, 2, 2],
                    "escala_turca_quinto_grado":   [2, 2, 1, 1, 2, 2, 2],
                    "escala_turca_sexto_grado":    [2, 1, 1, 2, 2, 2, 2],
                    "escala_turca_septimo_grado":  [1, 1, 2, 2, 2, 2, 2]},

  "escala_hungara_mayor":{
                    "escala_hungara_mayor_primer_grado":        [3, 1, 2, 1, 2, 1, 2],
                    "escala_hungara_mayor_segundo_grado":       [1, 2, 1, 2, 1, 2, 3],
                    "escala_hungara_mayor_tercer_grado":        [2, 1, 2, 1, 2, 3, 1],
                    "escala_hungara_mayor_cuarto_grado":        [1, 2, 1, 2, 3, 1, 2],
                    "escala_hungara_mayor_quinto_grado":        [2, 1, 2, 3, 1, 2, 1],
                    "escala_hungara_mayor_sexto_grado":         [1, 2, 3, 1, 2, 1, 2],
                    "escala_hungara_mayor_septimo_grado":       [2, 3, 1, 2, 1, 2, 1]},

  "escala_englimatic":{
                    "escala_englimatic_primer_grado":  [1, 3, 2, 2, 2, 1, 1],
                    "escala_englimatic_segundo_grado": [3, 2, 2, 2, 1, 1, 1],
                    "escala_englimatic_tercer_grado":  [1, 3, 2, 2, 2, 1, 1],
                    "escala_englimatic_cuarto_grado":  [2, 2, 1, 1, 1, 3, 2],
                    "escala_englimatic_quinto_grado":  [2, 1, 1, 1, 3, 2, 2],
                    "escala_englimatic_sexto_grado":   [1, 1, 1, 3, 2, 2, 2],
                    "escala_englimatic_septimo_grado": [1, 1, 1, 3, 2, 2, 2]},

  "escala_napolitana_menor":{
                    "escala_napolitana_menor_primer_grado":  [1, 2, 2, 2, 1, 3, 1],
                    "escala_lidia_#6":                       [2, 2, 2, 1, 3, 1, 1],
                    "escala_napolitana_menor_tercer_grado":  [2, 2, 1, 3, 1, 1, 2],
                    "escala_gipsy":                          [2, 1, 3, 1, 1, 2, 2],
                    "escala_locria_mayor":                   [1, 3, 1, 1, 2, 2, 2],
                    "escala_napolitana_menor_sexto_grado":   [3, 1, 1, 2, 2, 2, 1],
                    "escala_napolitana_menor_septimo_grado": [1, 1, 2, 2, 2, 1, 3]},

  "escala_napolitana_mayor":{
                    "escala_persa":                   [1, 3, 1, 2, 1, 3, 1],
                    "escala_lidia_#2_#6":             [3, 1, 2, 1, 3, 1, 1],
                    "escala_frigia_b4_bb7":           [1, 2, 1, 3, 1, 1, 3],
                    "escala_doble_armonica_menor":    [2, 1, 3, 1, 1, 3, 1],
                    "escala_oriental":                [1, 3, 1, 1, 3, 1, 2],
                    "escala_jonica_aumentada_#2":     [3, 1, 1, 3, 1, 2, 1],
                    "escala_locria_bb3_bb7":          [1, 1, 3, 1, 2, 1, 3]},

  "escala_de_hexatónica_blues":{
                    "escala_hexatonica_de_blues_primer_grado":  [3, 2, 1, 1, 3, 2],
                    "escala_hexatonica_de_blues_segundo_grado": [2, 1, 1, 3, 2, 3],
                    "escala_hexatonica_de_blues_tercer_grado":  [1, 1, 3, 2, 3, 2],
                    "escala_hexatonica_de_blues_cuarto_grado":  [1, 3, 2, 3, 2, 1],
                    "escala_hexatonica_de_blues_quinto_grado":  [3, 2, 3, 2, 1, 1],
                    "escala_hexatonica_de_blues_sexto_grado":   [2, 3, 2, 1, 1, 3]},

  "escala_heptatónica_de_blues":{
                    "escala_heptatonica_de_blues_primer_grado":  [3, 1, 1, 1, 1, 3, 2],
                    "escala_heptatonica_de_blues_segundo_grado": [1, 1, 1, 1, 3, 2, 3],
                    "escala_heptatonica_de_blues_tercer_grado":  [1, 1, 1, 3, 2, 3, 1],
                    "escala_heptatonica_de_blues_cuarto_grado":  [1, 1, 3, 2, 3, 1, 1],
                    "escala_heptatonica_de_blues_quinto_grado":  [1, 3, 2, 3, 1, 1, 1],
                    "escala_heptatonica_de_blues_sexto_grado":   [3, 2, 3, 1, 1, 1, 1],
                    "escala_heptatonica_de_blues_septimo_grado": [2, 3, 1, 1, 1, 1, 3]},

  "escalas_simétricas":{
                    "escala_tono_a_tono":             [2, 2, 2, 2, 2, 2],
                    "escala_semitono_a_tono":         [1, 2, 1, 2, 1, 2, 1, 2],
                    "escala_tono_a_semitono":         [2, 1, 2, 1, 2, 1, 2, 1],
                    "escala_cromatica":               [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]}}
`)

	return scaleByte
}

/*TestRepo is a*/
func TestRepo() []byte {
	tbyte := []byte(`
		{
		  "escala_natural":{
		                  "1" :       {"name" : ["escala_jonica"],
																	"formula" : ["2", "2", "1", "2", "2", "2", "1"],
																	"intervals" : ["2", "3", "4", "5", "6", "7"]},
		                  "2" :       {"name" : ["escala_dorica"],
																	"formula" : ["2", "1", "2", "2", "2", "1", "2"],
																	"intervals" : ["2", "b3", "4", "5", "6", "b7"]},
		                  "3" :       {"name" : ["escala_frigia"],
																	"formula" : ["1", "2", "2", "2", "1", "2", "2"],
																	"intervals" : ["b2", "b3", "4", "5", "b6", "b7"]},
		                  "4" :       {"name" : ["escala_lidia"],
																	"formula" : ["2", "2", "2", "1", "2", "2", "1"],
																	"intervals" : ["2", "3", "#4", "5", "6", "7"]},
		                  "5" :       {"name" : ["escala_mixolidia"],
																	"formula" : ["2", "2", "1", "2", "2", "1", "2"],
																	"intervals" : ["2", "3", "4", "5", "6", "b7"]},
		                  "6" :       {"name" : ["escala_eolica"],
																	"formula" : ["2", "1", "2", "2", "1", "2", "2"],
																	"intervals" : ["2", "b3", "4", "5", "b6", "b7"]},
		                  "7" :       {"name" : ["escala_locria"],
																	"formula" : ["1", "2", "2", "1", "2", "2", "2"],
																	"intervals" : ["b2", "b3", "4", "b5", "b6", "b7"]}}}

		`)
	return tbyte
}
