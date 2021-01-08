package google

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBigQueryTable_Basic(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableTimePartitioning(datasetID, tableID, "DAY"),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableUpdated(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_Kms(t *testing.T) {
	t.Parallel()
	resourceName := "google_bigquery_table.test"
	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	kms := BootstrapKMSKey(t)
	cryptoKeyName := kms.CryptoKey.Name

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableKms(cryptoKeyName, datasetID, tableID),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_HourlyTimePartitioning(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableTimePartitioning(datasetID, tableID, "HOUR"),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableUpdated(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_MonthlyTimePartitioning(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableTimePartitioning(datasetID, tableID, "MONTH"),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableUpdated(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_YearlyTimePartitioning(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableTimePartitioning(datasetID, tableID, "YEAR"),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableUpdated(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_HivePartitioning(t *testing.T) {
	t.Parallel()
	bucketName := testBucketName(t)
	resourceName := "google_bigquery_table.test"
	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableHivePartitioning(bucketName, datasetID, tableID),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_HivePartitioningCustomSchema(t *testing.T) {
	t.Parallel()
	bucketName := testBucketName(t)
	resourceName := "google_bigquery_table.test"
	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableHivePartitioningCustomSchema(bucketName, datasetID, tableID),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"external_data_configuration.0.schema"},
			},
		},
	})
}

func TestAccBigQueryTable_RangePartitioning(t *testing.T) {
	t.Parallel()
	resourceName := "google_bigquery_table.test"
	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableRangePartitioning(datasetID, tableID),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_View(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableWithView(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_updateView(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableWithView(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableWithNewSqlView(datasetID, tableID),
			},
			{
				ResourceName:      "google_bigquery_table.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTable_MaterializedView_DailyTimePartioning_Basic(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	materialized_viewID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	query := fmt.Sprintf("SELECT count(some_string) as count, some_int, ts FROM `%s.%s` WHERE DATE(ts) = '2019-01-01' GROUP BY some_int, ts", datasetID, tableID)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableWithMatViewDailyTimePartitioning_basic(datasetID, tableID, materialized_viewID, query),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				ResourceName:            "google_bigquery_table.mv_test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
		},
	})
}

func TestAccBigQueryTable_MaterializedView_DailyTimePartioning_Update(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	materialized_viewID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	query := fmt.Sprintf("SELECT count(some_string) as count, some_int, ts FROM `%s.%s` WHERE DATE(ts) = '2019-01-01' GROUP BY some_int, ts", datasetID, tableID)

	enable_refresh := "false"
	refresh_interval_ms := "3600000"

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableWithMatViewDailyTimePartitioning_basic(datasetID, tableID, materialized_viewID, query),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				ResourceName:            "google_bigquery_table.mv_test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				Config: testAccBigQueryTableWithMatViewDailyTimePartitioning(datasetID, tableID, materialized_viewID, enable_refresh, refresh_interval_ms, query),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				ResourceName:            "google_bigquery_table.mv_test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
		},
	})
}

func TestAccBigQueryExternalDataTable_CSV(t *testing.T) {
	t.Parallel()

	bucketName := testBucketName(t)
	objectName := fmt.Sprintf("tf_test_%s.csv", randString(t, 10))

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableFromGCS(datasetID, tableID, bucketName, objectName, TEST_CSV, "CSV", "\\\""),
				Check:  testAccCheckBigQueryExtData(t, "\""),
			},
			{
				Config: testAccBigQueryTableFromGCS(datasetID, tableID, bucketName, objectName, TEST_CSV, "CSV", ""),
				Check:  testAccCheckBigQueryExtData(t, ""),
			},
		},
	})
}

func TestAccBigQueryDataTable_sheet(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableFromSheet(context),
			},
			{
				ResourceName:      "google_bigquery_table.table",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryDataTable_jsonEquivalency(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTable_jsonEq(datasetID, tableID),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				Config: testAccBigQueryTable_jsonEqUpdate(datasetID, tableID),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
		},
	})
}

func TestAccBigQueryDataTable_schemaToFieldsAndBack(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTable_jsonEq(datasetID, tableID),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
			{
				Config: testAccBigQueryTable_fields(datasetID, tableID),
				Check:  testAccCheckBigQueryTableFields(t),
			},
			{
				Config: testAccBigQueryTable_jsonEq(datasetID, tableID),
			},
			{
				ResourceName:            "google_bigquery_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "last_modified_time"},
			},
		},
	})
}

func TestAccBigQueryDataTable_fields(t *testing.T) {
	t.Parallel()

	datasetID := fmt.Sprintf("tf_test_%s", randString(t, 10))
	tableID := fmt.Sprintf("tf_test_%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTable_fields(datasetID, tableID),
				Check:  testAccCheckBigQueryTableFields(t),
			},
		},
	})
}

func TestUnitBigQueryDataTable_jsonEquivalency(t *testing.T) {
	t.Parallel()

	for i, testcase := range testUnitBigQueryDataTableJSONEquivalencyTestCases {
		var a, b interface{}
		if err := json.Unmarshal([]byte(testcase.jsonA), &a); err != nil {
			panic(fmt.Sprintf("unable to unmarshal json - %v", err))
		}
		if err := json.Unmarshal([]byte(testcase.jsonB), &b); err != nil {
			panic(fmt.Sprintf("unable to unmarshal json - %v", err))
		}
		eq, err := jsonCompareWithMapKeyOverride(a, b, bigQueryTableMapKeyOverride)
		if err != nil {
			t.Errorf("ahhhh an error I did not expect this! especially not on testscase %v - %s", i, err)
		}
		if eq != testcase.equivalent {
			t.Errorf("expected equivalency result of %v but got %v for testcase number %v", testcase.equivalent, eq, i)
		}
	}
}

func TestUnitBigQueryDataTable_schemaIsChangable(t *testing.T) {
	t.Parallel()
	for _, testcase := range testUnitBigQueryDataTableIsChangableTestCases {
		testcase.check(t)
		testcaseNested := &testUnitBigQueryDataTableJSONChangeableTestCase{
			testcase.name + "Nested",
			fmt.Sprintf("[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"fields\" : %s }]", testcase.jsonOld),
			fmt.Sprintf("[{\"name\": \"someValue\", \"type\" : \"INT64\", \"fields\" : %s }]", testcase.jsonNew),
			testcase.changeable,
		}
		testcaseNested.check(t)
	}
}

func TestUnitBigQueryDataTable_customizedDiffSchema(t *testing.T) {
	t.Parallel()

	tcs := testUnitBigQueryDataCustomDiffFieldChangeTestcases

	for _, changeableTestcase := range testUnitBigQueryDataTableIsChangableTestCases {
		extraTestcase := testUnitBigQueryDataCustomDiffFieldChangeTestcase{
			name:           changeableTestcase.name,
			schemaBefore:   changeableTestcase.jsonOld,
			schemaAfter:    changeableTestcase.jsonNew,
			shouldForceNew: !changeableTestcase.changeable,
		}
		tcs = append(tcs, extraTestcase)
	}

	for _, testcase := range tcs {
		testcase.check(t)
	}
}

type testUnitBigQueryDataTableJSONEquivalencyTestCase struct {
	jsonA      string
	jsonB      string
	equivalent bool
}

type testUnitBigQueryDataTableJSONChangeableTestCase struct {
	name       string
	jsonOld    string
	jsonNew    string
	changeable bool
}

type testUnitBigQueryDataCustomDiffField struct {
	mode        string
	description string
}

type testUnitBigQueryDataCustomDiffFieldChangeTestcase struct {
	name           string
	fieldBefore    []testUnitBigQueryDataCustomDiffField
	fieldAfter     []testUnitBigQueryDataCustomDiffField
	schemaBefore   string
	schemaAfter    string
	shouldForceNew bool
}

func (testcase *testUnitBigQueryDataTableJSONChangeableTestCase) check(t *testing.T) {
	var old, new interface{}
	if err := json.Unmarshal([]byte(testcase.jsonOld), &old); err != nil {
		panic(fmt.Sprintf("unable to unmarshal json - %v", err))
	}
	if err := json.Unmarshal([]byte(testcase.jsonNew), &new); err != nil {
		panic(fmt.Sprintf("unable to unmarshal json - %v", err))
	}
	changeable, err := resourceBigQueryTableSchemaIsChangable(old, new)
	if err != nil {
		t.Errorf("ahhhh an error I did not expect this! especially not on testscase %s - %s", testcase.name, err)
	}
	if changeable != testcase.changeable {
		t.Errorf("expected changeable result of %v but got %v for testcase %s", testcase.changeable, changeable, testcase.name)
	}
}

func (testcase *testUnitBigQueryDataCustomDiffFieldChangeTestcase) check(t *testing.T) {
	d := &ResourceDiffMock{
		Before: map[string]interface{}{},
		After:  map[string]interface{}{},
	}
	if testcase.fieldBefore != nil {
		d.Before["field.#"] = len(testcase.fieldBefore)
		for i, f := range testcase.fieldBefore {
			d.Before[fmt.Sprintf("field.%d.mode", i)] = f.mode
			d.Before[fmt.Sprintf("field.%d.description", i)] = f.description
		}
	}
	if testcase.fieldAfter != nil {
		d.After["field.#"] = len(testcase.fieldAfter)
		for i, f := range testcase.fieldAfter {
			d.After[fmt.Sprintf("field.%d.mode", i)] = f.mode
			d.After[fmt.Sprintf("field.%d.description", i)] = f.description
		}
	}
	if testcase.schemaBefore != "" {
		d.Before["schema"] = testcase.schemaBefore
	}
	if testcase.schemaAfter != "" {
		d.After["schema"] = testcase.schemaAfter
	}

	err := resourceBigQueryTableSchemaCustomizeDiffFunc(d)
	if err != nil {
		t.Errorf("error on testcase %s - %w", testcase.name, err)
	}
	if testcase.shouldForceNew != d.IsForceNew {
		t.Errorf("%s: expected d.IsForceNew to be %v, but was %v", testcase.name, testcase.shouldForceNew, d.IsForceNew)
	}
}

var testUnitBigQueryDataTableIsChangableTestCases = []testUnitBigQueryDataTableJSONChangeableTestCase{
	{
		"defaultEquality",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		true,
	},
	{
		"arraySizeIncreases",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }, {\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		true,
	},
	{
		"arraySizeDecreases",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }, {\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		false,
	},
	{
		"descriptionChanges",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		true,
	},
	{
		"typeInteger",
		"[{\"name\": \"someValue\", \"type\" : \"INTEGER\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"INT64\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		true,
	},
	{
		"typeFloat",
		"[{\"name\": \"someValue\", \"type\" : \"FLOAT\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"FLOAT64\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		true,
	},
	{
		"typeBool",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"BOOL\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		true,
	},
	{
		"typeRandom",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"NULLABLE\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"DATETIME\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		false,
	},
	{
		"typeModeReqToNull",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"REQUIRED\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"NULLABLE\", \"description\" : \"some new value\" }]",
		true,
	},
	{
		"typeModeRandom",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"REQUIRED\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"REPEATED\", \"description\" : \"some new value\" }]",
		false,
	},
	{
		"typeModeOmission",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"mode\" : \"REQUIRED\", \"description\" : \"someVal\" }]",
		"[{\"name\": \"someValue\", \"type\" : \"BOOLEAN\", \"description\" : \"some new value\" }]",
		false,
	},
}

var testUnitBigQueryDataCustomDiffFieldChangeTestcases = []testUnitBigQueryDataCustomDiffFieldChangeTestcase{
	{
		name: "control",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "Nullable",
			description: "someValue",
		}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "Nullable",
			description: "someValue",
		}},
		shouldForceNew: false,
	},
	{
		name: "requiredToNullable",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "NULLABLE",
			description: "someValue",
		}},
		shouldForceNew: false,
	},
	{
		name: "descriptionChanged",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "some other value",
		}},
		shouldForceNew: false,
	},
	{
		name: "nullToRequired",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "NULLABLE",
			description: "someValue",
		}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		}},
		shouldForceNew: true,
	},
	{
		name: "arraySizeIncreases",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		},
			{
				mode:        "REQUIRED",
				description: "some other value",
			}},
		shouldForceNew: false,
	},
	{
		name: "arraySizeDecreases",
		fieldBefore: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		},
			{
				mode:        "REQUIRED",
				description: "someValue",
			}},
		fieldAfter: []testUnitBigQueryDataCustomDiffField{{
			mode:        "REQUIRED",
			description: "someValue",
		}},
		shouldForceNew: true,
	},
}

var testUnitBigQueryDataTableJSONEquivalencyTestCases = []testUnitBigQueryDataTableJSONEquivalencyTestCase{
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"finalKey\" : {} }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"finalKey\" : {} }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"finalKey\" : {} }]",
		"[{\"someKey\": \"someValue\", \"finalKey\" : {} }]",
		false,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"mode\": \"NULLABLE\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\" }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"mode\": \"NULLABLE\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"mode\": \"somethingRandom\"  }]",
		false,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"description\": \"\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\" }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"description\": \"\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"description\": \"somethingRandom\"  }]",
		false,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"INTEGER\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"INT64\"  }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"INTEGER\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"somethingRandom\"  }]",
		false,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"FLOAT\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"FLOAT64\"  }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"FLOAT\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\" }]",
		false,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"BOOLEAN\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"BOOL\"  }]",
		true,
	},
	{
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\", \"type\": \"BOOLEAN\"  }]",
		"[{\"someKey\": \"someValue\", \"anotherKey\" : \"anotherValue\" }]",
		false,
	},
	{
		"[1,2,3]",
		"[1,2,3]",
		true,
	},
	{
		"[1,2,3]",
		"[1,2,\"banana\"]",
		false,
	},
}

func testAccCheckBigQueryExtData(t *testing.T, expectedQuoteChar string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_table" {
				continue
			}

			config := googleProviderConfig(t)
			dataset := rs.Primary.Attributes["dataset_id"]
			table := rs.Primary.Attributes["table_id"]
			res, err := config.NewBigQueryClient(config.userAgent).Tables.Get(config.Project, dataset, table).Do()
			if err != nil {
				return err
			}

			if res.Type != "EXTERNAL" {
				return fmt.Errorf("Table \"%s.%s\" is of type \"%s\", expected EXTERNAL.", dataset, table, res.Type)
			}
			edc := res.ExternalDataConfiguration
			cvsOpts := edc.CsvOptions
			if cvsOpts == nil || *cvsOpts.Quote != expectedQuoteChar {
				return fmt.Errorf("Table \"%s.%s\" quote should be '%s' but was '%s'", dataset, table, expectedQuoteChar, *cvsOpts.Quote)
			}
		}
		return nil
	}
}

func testAccCheckBigQueryTableDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_table" {
				continue
			}

			config := googleProviderConfig(t)
			_, err := config.NewBigQueryClient(config.userAgent).Tables.Get(config.Project, rs.Primary.Attributes["dataset_id"], rs.Primary.Attributes["table_id"]).Do()
			if err == nil {
				return fmt.Errorf("Table still present")
			}
		}

		return nil
	}
}

func testAccCheckBigQueryTableFields(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_table" {
				continue
			}
			config := googleProviderConfig(t)
			bqt, err := config.NewBigQueryClient(config.userAgent).Tables.Get(config.Project, rs.Primary.Attributes["dataset_id"], rs.Primary.Attributes["table_id"]).Do()
			if err != nil {
				return err
			}
			if bqt.Schema.Fields[1].Description != "testegg" {
				t.Errorf("schema field not set to testegg for fields scenario")
			}
		}
		return nil
	}
}

func testAccBigQueryTableTimePartitioning(datasetID, tableID, partitioningType string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type                     = "%s"
    field                    = "ts"
    require_partition_filter = true
  }
  clustering = ["some_int", "some_string"]
  schema     = <<EOH
[
  {
    "name": "ts",
    "type": "TIMESTAMP"
  },
  {
    "name": "some_string",
    "type": "STRING"
  },
  {
    "name": "some_int",
    "type": "INTEGER"
  },
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
    "name": "lon",
    "type": "FLOAT"
    }
    ]
  }
    ]
  }
]
EOH

}
`, datasetID, tableID, partitioningType)
}

func testAccBigQueryTableKms(cryptoKeyName, datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
    dataset_id = "%s"
}

data "google_bigquery_default_service_account" "acct" {}

resource "google_kms_crypto_key_iam_member" "allow" {
  crypto_key_id = "%s"
  role = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member = "serviceAccount:${data.google_bigquery_default_service_account.acct.email}"
  depends_on = ["google_bigquery_dataset.test"]
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = "${google_bigquery_dataset.test.dataset_id}"

  time_partitioning {
    type = "DAY"
    field = "ts"
  }

  encryption_configuration {
    kms_key_name = "${google_kms_crypto_key_iam_member.allow.crypto_key_id}"
  }

  schema = <<EOH
[
  {
    "name": "ts",
    "type": "TIMESTAMP"
  },
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
    "name": "lon",
    "type": "FLOAT"
    }
    ]
  }
    ]
  }
]
EOH
}
`, datasetID, cryptoKeyName, tableID)
}

func testAccBigQueryTableHivePartitioning(bucketName, datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "test" {
  name          = "%s"
  force_destroy = true
}

resource "google_storage_bucket_object" "test" {
  name    = "key1=20200330/init.csv"
  content = ";"
  bucket  = google_storage_bucket.test.name
}

resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  external_data_configuration {
    source_format = "CSV"
    autodetect = true
    source_uris= ["gs://${google_storage_bucket.test.name}/*"]

    hive_partitioning_options {
      mode = "AUTO"
      source_uri_prefix = "gs://${google_storage_bucket.test.name}/"
    }

  }
  depends_on = ["google_storage_bucket_object.test"]
}
`, bucketName, datasetID, tableID)
}

func testAccBigQueryTableHivePartitioningCustomSchema(bucketName, datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "test" {
  name          = "%s"
  force_destroy = true
}

resource "google_storage_bucket_object" "test" {
  name    = "key1=20200330/data.json"
  content = "{\"name\":\"test\", \"last_modification\":\"2020-04-01\"}"
  bucket  = google_storage_bucket.test.name
}

resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  external_data_configuration {
    source_format = "NEWLINE_DELIMITED_JSON"
    autodetect = false
    source_uris= ["gs://${google_storage_bucket.test.name}/*"]

    hive_partitioning_options {
      mode = "CUSTOM"
      source_uri_prefix = "gs://${google_storage_bucket.test.name}/{key1:STRING}"
    }

    schema = <<EOH
[
  {
    "name": "name",
    "type": "STRING"
  },
  {
    "name": "last_modification",
    "type": "DATE"
  }
]
EOH
        }
  depends_on = ["google_storage_bucket_object.test"]
}
`, bucketName, datasetID, tableID)
}

func testAccBigQueryTableRangePartitioning(datasetID, tableID string) string {
	return fmt.Sprintf(`
  resource "google_bigquery_dataset" "test" {
    dataset_id = "%s"
  }

  resource "google_bigquery_table" "test" {
    table_id   = "%s"
    dataset_id = google_bigquery_dataset.test.dataset_id

    range_partitioning {
      field = "id"
      range {
        start    = 0
        end      = 10000
        interval = 100
      }
    }

    schema = <<EOH
[
  {
    "name": "ts",
    "type": "TIMESTAMP"
  },
  {
    "name": "id",
    "type": "INTEGER"
  }
]
EOH
}
  `, datasetID, tableID)
}

func testAccBigQueryTableWithView(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type = "DAY"
  }

  view {
    query          = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]"
    use_legacy_sql = true
  }
}
`, datasetID, tableID)
}

func testAccBigQueryTableWithNewSqlView(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type = "DAY"
  }

  view {
    query          = "%s"
    use_legacy_sql = false
  }
}
`, datasetID, tableID, "SELECT state FROM `lookerdata.cdc.project_tycho_reports`")
}

func testAccBigQueryTableWithMatViewDailyTimePartitioning_basic(datasetID, tableID, mViewID, query string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type                     = "DAY"
    field                    = "ts"
    require_partition_filter = true
  }
  clustering = ["some_int", "some_string"]
  schema     = <<EOH
[
  {
    "name": "ts",
    "type": "TIMESTAMP"
  },
  {
    "name": "some_string",
    "type": "STRING"
  },
  {
    "name": "some_int",
    "type": "INTEGER"
  },
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
    "name": "lon",
    "type": "FLOAT"
    }
    ]
  }
    ]
  }
]
EOH

}

resource "google_bigquery_table" "mv_test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type    = "DAY"
    field   = "ts"
  }

  materialized_view {
    query          = "%s"
  }

  depends_on = [
    google_bigquery_table.test,
  ]
}
`, datasetID, tableID, mViewID, query)
}

func testAccBigQueryTableWithMatViewDailyTimePartitioning(datasetID, tableID, mViewID, enable_refresh, refresh_interval, query string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type                     = "DAY"
    field                    = "ts"
    require_partition_filter = true
  }
  clustering = ["some_int", "some_string"]
  schema     = <<EOH
[
  {
    "name": "ts",
    "type": "TIMESTAMP"
  },
  {
    "name": "some_string",
    "type": "STRING"
  },
  {
    "name": "some_int",
    "type": "INTEGER"
  },
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
    "name": "lon",
    "type": "FLOAT"
    }
    ]
  }
    ]
  }
]
EOH

}

resource "google_bigquery_table" "mv_test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type    = "DAY"
    field   = "ts"
  }

  materialized_view {
    enable_refresh = "%s"
    refresh_interval_ms = "%s"
    query          = "%s"
  }

  depends_on = [
    google_bigquery_table.test,
  ]
}
`, datasetID, tableID, mViewID, enable_refresh, refresh_interval, query)
}

func testAccBigQueryTableUpdated(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  time_partitioning {
    type = "DAY"
  }

  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH

}
`, datasetID, tableID)
}

func testAccBigQueryTableFromGCS(datasetID, tableID, bucketName, objectName, content, format, quoteChar string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_storage_bucket" "test" {
  name          = "%s"
  force_destroy = true
}

resource "google_storage_bucket_object" "test" {
  name    = "%s"
  content = <<EOF
%s
EOF

  bucket = google_storage_bucket.test.name
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id
  external_data_configuration {
    autodetect    = true
    source_format = "%s"
    csv_options {
      encoding = "UTF-8"
      quote    = "%s"
    }

    source_uris = [
      "gs://${google_storage_bucket.test.name}/${google_storage_bucket_object.test.name}",
    ]
  }
}
`, datasetID, bucketName, objectName, content, tableID, format, quoteChar)
}

func testAccBigQueryTableFromSheet(context map[string]interface{}) string {
	return Nprintf(`
  resource "google_bigquery_table" "table" {
    dataset_id = google_bigquery_dataset.dataset.dataset_id
    table_id   = "tf_test_sheet_%{random_suffix}"

    external_data_configuration {
      autodetect            = true
      source_format         = "GOOGLE_SHEETS"
      ignore_unknown_values = true

      google_sheets_options {
      skip_leading_rows = 1
      }

      source_uris = [
      "https://drive.google.com/open?id=xxxx",
      ]
    }

    schema = <<EOF
    [
    {
      "name": "permalink",
      "type": "STRING",
      "mode": "NULLABLE",
      "description": "The Permalink"
    },
    {
      "name": "state",
      "type": "STRING",
      "mode": "NULLABLE",
      "description": "State where the head office is located"
    }
    ]
    EOF
    }

    resource "google_bigquery_dataset" "dataset" {
    dataset_id                  = "tf_test_ds_%{random_suffix}"
    friendly_name               = "test"
    description                 = "This is a test description"
    location                    = "EU"
    default_table_expiration_ms = 3600000

    labels = {
      env = "default"
    }
    }
`, context)
}

func testAccBigQueryTable_jsonEq(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  friendly_name = "bigquerytest"
  labels = {
    "terrafrom_managed" = "true"
  }

  schema = jsonencode(
    [
      {
        description = "Time snapshot was taken, in Epoch milliseconds. Same across all rows and all tables in the snapshot, and uniquely defines a particular snapshot."
        name        = "snapshot_timestamp"
        mode        = "NULLABLE"
        type        = "INTEGER"
      },
      {
        description = "Timestamp of dataset creation"
        name        = "creation_time"
        type        = "TIMESTAMP"
      },
    ])
}
`, datasetID, tableID)
}

func testAccBigQueryTable_jsonEqUpdate(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  friendly_name = "bigquerytest"
  labels = {
    "terrafrom_managed" = "true"
  }

  schema = jsonencode(
    [
      {
        description = "Time snapshot was taken, in Epoch milliseconds. Same across all rows and all tables in the snapshot, and uniquely defines a particular snapshot."
        name        = "snapshot_timestamp"
        type        = "INTEGER"
      },
      {
        description = "Timestamp of dataset creation"
        name        = "creation_time"
        type        = "TIMESTAMP"
      },
    ])
}
`, datasetID, tableID)
}

func testAccBigQueryTable_fields(datasetID, tableID string) string {
	return fmt.Sprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "%s"
}

resource "google_bigquery_table" "test" {
  table_id   = "%s"
  dataset_id = google_bigquery_dataset.test.dataset_id

  friendly_name = "bigquerytest"
  labels = {
    "terrafrom_managed" = "true"
	}

	field {
		description = "Time snapshot was taken, in Epoch milliseconds. Same across all rows and all tables in the snapshot, and uniquely defines a particular snapshot."
		name        = "snapshot_timestamp"
		mode        = "NULLABLE"
		type        = "INTEGER"
	}

	field {
		description = "testegg"
		name        = "creation_time"
		type        = "TIMESTAMP"
	}
}
`, datasetID, tableID)
}

var TEST_CSV = `lifelock,LifeLock,,web,Tempe,AZ,1-May-07,6850000,USD,b
lifelock,LifeLock,,web,Tempe,AZ,1-Oct-06,6000000,USD,a
lifelock,LifeLock,,web,Tempe,AZ,1-Jan-08,25000000,USD,c
mycityfaces,MyCityFaces,7,web,Scottsdale,AZ,1-Jan-08,50000,USD,seed
flypaper,Flypaper,,web,Phoenix,AZ,1-Feb-08,3000000,USD,a
infusionsoft,Infusionsoft,105,software,Gilbert,AZ,1-Oct-07,9000000,USD,a
gauto,gAuto,4,web,Scottsdale,AZ,1-Jan-08,250000,USD,seed
chosenlist-com,ChosenList.com,5,web,Scottsdale,AZ,1-Oct-06,140000,USD,seed
chosenlist-com,ChosenList.com,5,web,Scottsdale,AZ,25-Jan-08,233750,USD,angel
`
