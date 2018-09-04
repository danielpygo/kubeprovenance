package provenance

import (
	"reflect"
	"strings"
	"testing"
)

func TestCollectProvenance(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			CollectProvenance()
		})
	}
}

func Test_readKindCompositionFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			readKindCompositionFile()
		})
	}
}

func TestNewProvenanceOfObject(t *testing.T) {
	tests := []struct {
		name string
		want *ProvenanceOfObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProvenanceOfObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProvenanceOfObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpec(t *testing.T) {
	tests := []struct {
		name string
		want *Spec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindProvenanceObjectByName(t *testing.T) {
	type args struct {
		name       string
		allObjects []ProvenanceOfObject
	}
	tests := []struct {
		name string
		args args
		want *ProvenanceOfObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindProvenanceObjectByName(tt.args.name, tt.args.allObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindProvenanceObjectByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_String(t *testing.T) {
	type fields struct {
		AttributeToData map[string]interface{}
		Version         int
		Timestamp       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spec{
				AttributeToData: tt.fields.AttributeToData,
				Version:         tt.fields.Version,
				Timestamp:       tt.fields.Timestamp,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Spec.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_String(t *testing.T) {
	tests := []struct {
		name string
		o    ObjectLineage
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("ObjectLineage.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSpecsInOrder(t *testing.T) {
	type args struct {
		o ObjectLineage
	}
	tests := []struct {
		name string
		args args
		want []Spec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSpecsInOrder(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSpecsInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_GetVersions(t *testing.T) {
	tests := []struct {
		name string
		o    ObjectLineage
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetVersions(); got != tt.want {
				t.Errorf("ObjectLineage.GetVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_SpecHistory(t *testing.T) {
	tests := []struct {
		name string
		o    ObjectLineage
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SpecHistory(); got != tt.want {
				t.Errorf("ObjectLineage.SpecHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_SpecHistoryInterval(t *testing.T) {
	type args struct {
		vNumStart int
		vNumEnd   int
	}
	tests := []struct {
		name string
		o    ObjectLineage
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SpecHistoryInterval(tt.args.vNumStart, tt.args.vNumEnd); got != tt.want {
				t.Errorf("ObjectLineage.SpecHistoryInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildAttributeRelationships(t *testing.T) {
	type args struct {
		specs         []Spec
		allQueryPairs [][]string
	}
	tests := []struct {
		name string
		args args
		want map[string][][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAttributeRelationships(tt.args.specs, tt.args.allQueryPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildAttributeRelationships() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildQueryPairsSlice(t *testing.T) {
	type args struct {
		queryArgMap map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildQueryPairsSlice(tt.args.queryArgMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildQueryPairsSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildQueryPairsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_Bisect(t *testing.T) {
	type args struct {
		argMap map[string]string
	}
	tests := []struct {
		name string
		o    ObjectLineage
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Bisect(tt.args.argMap); got != tt.want {
				t.Errorf("ObjectLineage.Bisect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleTrivialFields(t *testing.T) {
	type args struct {
		qkey      string
		qval      string
		mkey      string
		fieldData string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleTrivialFields(tt.args.qkey, tt.args.qval, tt.args.mkey, tt.args.fieldData); got != tt.want {
				t.Errorf("handleTrivialFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleSimpleFields(t *testing.T) {
	type args struct {
		vStringSlice []string
		qkey         string
		qval         string
		mkey         string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleSimpleFields(tt.args.vStringSlice, tt.args.qkey, tt.args.qval, tt.args.mkey); got != tt.want {
				t.Errorf("handleSimpleFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleCompositeFields(t *testing.T) {
	type args struct {
		vSliceMap        []map[string]string
		mapRelationships map[string][][]string
		qkey             string
		qval             string
		mkey             string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleCompositeFields(tt.args.vSliceMap, tt.args.mapRelationships, tt.args.qkey, tt.args.qval, tt.args.mkey); got != tt.want {
				t.Errorf("handleCompositeFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_all(t *testing.T) {
	type args struct {
		boolSlice []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := all(tt.args.boolSlice); got != tt.want {
				t.Errorf("all() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareMaps(t *testing.T) {
	type args struct {
		mapSlice1 []map[string]string
		mapSlice2 []map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareMaps(tt.args.mapSlice1, tt.args.mapSlice2); got != tt.want {
				t.Errorf("compareMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_FullDiff(t *testing.T) {
	type args struct {
		vNumStart int
		vNumEnd   int
	}
	tests := []struct {
		name string
		o    ObjectLineage
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.FullDiff(tt.args.vNumStart, tt.args.vNumEnd); got != tt.want {
				t.Errorf("ObjectLineage.FullDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiff(t *testing.T) {
	type args struct {
		b         *strings.Builder
		fieldName string
		data1     interface{}
		data2     interface{}
		vNumStart int
		vNumEnd   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiff(tt.args.b, tt.args.fieldName, tt.args.data1, tt.args.data2, tt.args.vNumStart, tt.args.vNumEnd); got != tt.want {
				t.Errorf("getDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectLineage_FieldDiff(t *testing.T) {
	type args struct {
		fieldName string
		vNumStart int
		vNumEnd   int
	}
	tests := []struct {
		name string
		o    ObjectLineage
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.FieldDiff(tt.args.fieldName, tt.args.vNumStart, tt.args.vNumEnd); got != tt.want {
				t.Errorf("ObjectLineage.FieldDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse()
		})
	}
}

func Test_parseRequestObject(t *testing.T) {
	type args struct {
		objectProvenance *ProvenanceOfObject
		requestObjBytes  []byte
		timestamp        string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseRequestObject(tt.args.objectProvenance, tt.args.requestObjBytes, tt.args.timestamp)
		})
	}
}

func Test_buildSpec(t *testing.T) {
	type args struct {
		spec map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Spec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildSpec(tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printMaps(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			printMaps()
		})
	}
}

func Test_getResourceKinds(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResourceKinds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResourceKinds() = %v, want %v", got, tt.want)
			}
		})
	}
}
