// Copyright 2013 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package session

import (
	"github.com/hanchuanchuan/goInception/mysql"
)

// 数据库类型
const (
	DBTypeMysql = iota
	DBTypeMariaDB
	DBTypeTiDB
)

// 审核阶段
const (
	StageOK byte = iota
	StageCheck
	StageExec
	StageBackup
)

// 审核状态
const (
	StatusAuditOk byte = iota
	StatusExecFail
	StatusExecOK
	StatusBackupFail
	StatusBackupOK
)

var (
	StageList  = [4]string{"RERUN", "CHECKED", "EXECUTED", "BACKUP"}
	StatusList = [5]string{"Audit Completed", "Execute failed", "Execute Successfully",
		"Execute Successfully\nBackup failed", "Execute Successfully\nBackup Successfully"}
)

// int类型map(用以数值列类型变更审核)
var IntegerOrderedMaps = map[string]int{
	"bit":       1,
	"tinyint":   2,
	"smallint":  3,
	"mediumint": 4,
	"int":       5,
	"bigint":    6,
}

// int类型map(用以数值列类型变更审核)
var IntegerOrderedMaps2 = map[byte]int{
	mysql.TypeBit:      1,
	mysql.TypeTiny:     2,
	mysql.TypeShort:    3,
	mysql.TypeInt24:    4,
	mysql.TypeLong:     5,
	mysql.TypeLonglong: 6,
}

// Keywords 数据库关键字
var Keywords = map[string]bool{
	"ACTION":                   true,
	"ADD":                      true,
	"ADDDATE":                  true,
	"ADMIN":                    true,
	"AFTER":                    true,
	"ALL":                      true,
	"ALGORITHM":                true,
	"ALTER":                    true,
	"ALWAYS":                   true,
	"ANALYZE":                  true,
	"AND":                      true,
	"ANY":                      true,
	"AS":                       true,
	"ASC":                      true,
	"ASCII":                    true,
	"AUTO_INCREMENT":           true,
	"AVG":                      true,
	"AVG_ROW_LENGTH":           true,
	"BEGIN":                    true,
	"BETWEEN":                  true,
	"BIGINT":                   true,
	"BINARY":                   true,
	"BINLOG":                   true,
	"BIT":                      true,
	"BIT_AND":                  true,
	"BIT_OR":                   true,
	"BIT_XOR":                  true,
	"BLOB":                     true,
	"BOOL":                     true,
	"BOOLEAN":                  true,
	"BOTH":                     true,
	"BTREE":                    true,
	"BUCKETS":                  true,
	"BY":                       true,
	"BYTE":                     true,
	"CANCEL":                   true,
	"CASCADE":                  true,
	"CASCADED":                 true,
	"CASE":                     true,
	"CAST":                     true,
	"CHANGE":                   true,
	"CHAR":                     true,
	"CHARACTER":                true,
	"CHARSET":                  true,
	"CHECK":                    true,
	"CHECKSUM":                 true,
	"CLEANUP":                  true,
	"CLIENT":                   true,
	"COALESCE":                 true,
	"COLLATE":                  true,
	"COLLATION":                true,
	"COLUMN":                   true,
	"COLUMNS":                  true,
	"COMMENT":                  true,
	"COMMIT":                   true,
	"INCEPTION":                true,
	"INCEPTION_MAGIC_START":    true,
	"INCEPTION_MAGIC_COMMIT":   true,
	"COMMITTED":                true,
	"COMPACT":                  true,
	"COMPRESSED":               true,
	"COMPRESSION":              true,
	"CONNECTION":               true,
	"CONSISTENT":               true,
	"CONSTRAINT":               true,
	"CONVERT":                  true,
	"COPY":                     true,
	"COUNT":                    true,
	"CREATE":                   true,
	"CROSS":                    true,
	"CURRENT_DATE":             true,
	"CURRENT_TIME":             true,
	"CURRENT_TIMESTAMP":        true,
	"CURRENT_USER":             true,
	"CURTIME":                  true,
	"DATA":                     true,
	"DATABASE":                 true,
	"DATABASES":                true,
	"DATE":                     true,
	"DATE_ADD":                 true,
	"DATE_SUB":                 true,
	"DATETIME":                 true,
	"DAY":                      true,
	"DAY_HOUR":                 true,
	"DAY_MICROSECOND":          true,
	"DAY_MINUTE":               true,
	"DAY_SECOND":               true,
	"DDL":                      true,
	"DEALLOCATE":               true,
	"DEC":                      true,
	"DECIMAL":                  true,
	"DEFAULT":                  true,
	"DEFINER":                  true,
	"DELAY_KEY_WRITE":          true,
	"DELAYED":                  true,
	"DELETE":                   true,
	"DESC":                     true,
	"DESCRIBE":                 true,
	"DISABLE":                  true,
	"DISTINCT":                 true,
	"DISTINCTROW":              true,
	"DIV":                      true,
	"DO":                       true,
	"DOUBLE":                   true,
	"DROP":                     true,
	"DUAL":                     true,
	"DUPLICATE":                true,
	"DYNAMIC":                  true,
	"ELSE":                     true,
	"ENABLE":                   true,
	"ENCLOSED":                 true,
	"END":                      true,
	"ENGINE":                   true,
	"ENGINES":                  true,
	"ENUM":                     true,
	"ESCAPE":                   true,
	"ESCAPED":                  true,
	"EVENT":                    true,
	"EVENTS":                   true,
	"EXCLUSIVE":                true,
	"EXECUTE":                  true,
	"EXISTS":                   true,
	"EXPLAIN":                  true,
	"EXTRACT":                  true,
	"FALSE":                    true,
	"FIELDS":                   true,
	"FIRST":                    true,
	"FIXED":                    true,
	"FLOAT":                    true,
	"FLUSH":                    true,
	"FOR":                      true,
	"FORCE":                    true,
	"FOREIGN":                  true,
	"FORMAT":                   true,
	"FROM":                     true,
	"FULL":                     true,
	"FULLTEXT":                 true,
	"FUNCTION":                 true,
	"GENERATED":                true,
	"GET_FORMAT":               true,
	"GLOBAL":                   true,
	"GRANT":                    true,
	"GRANTS":                   true,
	"GROUP":                    true,
	"GROUP_CONCAT":             true,
	"HASH":                     true,
	"HAVING":                   true,
	"HIGH_PRIORITY":            true,
	"HOUR":                     true,
	"HOUR_MICROSECOND":         true,
	"HOUR_MINUTE":              true,
	"HOUR_SECOND":              true,
	"IDENTIFIED":               true,
	"IF":                       true,
	"IGNORE":                   true,
	"IN":                       true,
	"INDEX":                    true,
	"INDEXES":                  true,
	"INFILE":                   true,
	"INNER":                    true,
	"INPLACE":                  true,
	"INSERT":                   true,
	"INT":                      true,
	"INT1":                     true,
	"INT2":                     true,
	"INT3":                     true,
	"INT4":                     true,
	"INT8":                     true,
	"INTEGER":                  true,
	"INTERVAL":                 true,
	"INTERNAL":                 true,
	"INTO":                     true,
	"INVOKER":                  true,
	"IS":                       true,
	"ISOLATION":                true,
	"JOBS":                     true,
	"JOB":                      true,
	"JOIN":                     true,
	"JSON":                     true,
	"KEY":                      true,
	"KEY_BLOCK_SIZE":           true,
	"KEYS":                     true,
	"KILL":                     true,
	"LEADING":                  true,
	"LEFT":                     true,
	"LESS":                     true,
	"LEVEL":                    true,
	"LIKE":                     true,
	"LIMIT":                    true,
	"LINES":                    true,
	"LOAD":                     true,
	"LOCAL":                    true,
	"LOCALTIME":                true,
	"LOCALTIMESTAMP":           true,
	"LOCK":                     true,
	"LONG":                     true,
	"LONGBLOB":                 true,
	"LONGTEXT":                 true,
	"LOW_PRIORITY":             true,
	"MASTER":                   true,
	"MAX":                      true,
	"MAX_CONNECTIONS_PER_HOUR": true,
	"MAX_EXECUTION_TIME":       true,
	"MAX_QUERIES_PER_HOUR":     true,
	"MAX_ROWS":                 true,
	"MAX_UPDATES_PER_HOUR":     true,
	"MAX_USER_CONNECTIONS":     true,
	"MAXVALUE":                 true,
	"MEDIUMBLOB":               true,
	"MEDIUMINT":                true,
	"MEDIUMTEXT":               true,
	"MERGE":                    true,
	"MICROSECOND":              true,
	"MIN":                      true,
	"MIN_ROWS":                 true,
	"MINUTE":                   true,
	"MINUTE_MICROSECOND":       true,
	"MINUTE_SECOND":            true,
	"MOD":                      true,
	"MODE":                     true,
	"MODIFY":                   true,
	"MONTH":                    true,
	"NAMES":                    true,
	"NATIONAL":                 true,
	"NATURAL":                  true,
	"NO":                       true,
	"NO_WRITE_TO_BINLOG":       true,
	"NONE":                     true,
	"NOT":                      true,
	"NOW":                      true,
	"NULL":                     true,
	"NUMERIC":                  true,
	"NVARCHAR":                 true,
	"OFFSET":                   true,
	"ON":                       true,
	"ONLY":                     true,
	"OPTION":                   true,
	"OR":                       true,
	"ORDER":                    true,
	"OUTER":                    true,
	"PACK_KEYS":                true,
	"PARTITION":                true,
	"PARTITIONS":               true,
	"PASSWORD":                 true,
	"PLUGINS":                  true,
	"POSITION":                 true,
	"PRECISION":                true,
	"PREPARE":                  true,
	"PRIMARY":                  true,
	"PRIVILEGES":               true,
	"PROCEDURE":                true,
	"PROCESS":                  true,
	"PROCESSLIST":              true,
	"PROFILES":                 true,
	"QUARTER":                  true,
	"QUERY":                    true,
	"QUERIES":                  true,
	"QUICK":                    true,
	"SHARD_ROW_ID_BITS":        true,
	"RANGE":                    true,
	"RECOVER":                  true,
	"READ":                     true,
	"REAL":                     true,
	"RECENT":                   true,
	"REDUNDANT":                true,
	"REFERENCES":               true,
	"REGEXP":                   true,
	"RELOAD":                   true,
	"RENAME":                   true,
	"REPEAT":                   true,
	"REPEATABLE":               true,
	"REPLACE":                  true,
	"REPLICATION":              true,
	"RESTRICT":                 true,
	"REVERSE":                  true,
	"REVOKE":                   true,
	"RIGHT":                    true,
	"RLIKE":                    true,
	"ROLLBACK":                 true,
	"ROUTINE":                  true,
	"ROW":                      true,
	"ROW_COUNT":                true,
	"ROW_FORMAT":               true,
	"SCHEMA":                   true,
	"SCHEMAS":                  true,
	"SECOND":                   true,
	"SECOND_MICROSECOND":       true,
	"SECURITY":                 true,
	"SELECT":                   true,
	"SERIALIZABLE":             true,
	"SESSION":                  true,
	"SET":                      true,
	"SEPARATOR":                true,
	"SHARE":                    true,
	"SHARED":                   true,
	"SHOW":                     true,
	"SIGNED":                   true,
	"SLAVE":                    true,
	"SLOW":                     true,
	"SMALLINT":                 true,
	"SNAPSHOT":                 true,
	"SOME":                     true,
	"SQL":                      true,
	"SQL_CACHE":                true,
	"SQL_CALC_FOUND_ROWS":      true,
	"SQL_NO_CACHE":             true,
	"START":                    true,
	"STARTING":                 true,
	"STATS":                    true,
	"STATS_BUCKETS":            true,
	"STATS_HISTOGRAMS":         true,
	"STATS_HEALTHY":            true,
	"STATS_META":               true,
	"STATS_PERSISTENT":         true,
	"STATUS":                   true,
	"STORED":                   true,
	"STRAIGHT_JOIN":            true,
	"SUBDATE":                  true,
	"SUBPARTITION":             true,
	"SUBPARTITIONS":            true,
	"SUBSTR":                   true,
	"SUBSTRING":                true,
	"SUM":                      true,
	"SUPER":                    true,
	"TABLE":                    true,
	"TABLES":                   true,
	"TABLESPACE":               true,
	"TEMPORARY":                true,
	"TEMPTABLE":                true,
	"TERMINATED":               true,
	"TEXT":                     true,
	"THAN":                     true,
	"THEN":                     true,
	"TIDB":                     true,
	"TIDB_HJ":                  true,
	"TIDB_INLJ":                true,
	"TIDB_SMJ":                 true,
	"TIME":                     true,
	"TIMESTAMP":                true,
	"TIMESTAMPADD":             true,
	"TIMESTAMPDIFF":            true,
	"TINYBLOB":                 true,
	"TINYINT":                  true,
	"TINYTEXT":                 true,
	"TO":                       true,
	"TOP":                      true,
	"TRACE":                    true,
	"TRAILING":                 true,
	"TRANSACTION":              true,
	"TRIGGER":                  true,
	"TRIGGERS":                 true,
	"TRIM":                     true,
	"TRUE":                     true,
	"TRUNCATE":                 true,
	"UNCOMMITTED":              true,
	"UNDEFINED":                true,
	"UNION":                    true,
	"UNIQUE":                   true,
	"UNKNOWN":                  true,
	"UNLOCK":                   true,
	"UNSIGNED":                 true,
	"UPDATE":                   true,
	"USAGE":                    true,
	"USE":                      true,
	"USER":                     true,
	"USING":                    true,
	"UTC_DATE":                 true,
	"UTC_TIME":                 true,
	"UTC_TIMESTAMP":            true,
	"VALUE":                    true,
	"VALUES":                   true,
	"VARBINARY":                true,
	"VARCHAR":                  true,
	"VARIABLES":                true,
	"VIEW":                     true,
	"VIRTUAL":                  true,
	"WARNINGS":                 true,
	"ERRORS":                   true,
	"WEEK":                     true,
	"WHEN":                     true,
	"WHERE":                    true,
	"WITH":                     true,
	"WRITE":                    true,
	"XOR":                      true,
	"YEAR":                     true,
	"YEAR_MONTH":               true,
	"ZEROFILL":                 true,
}
