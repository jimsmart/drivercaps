
# Driver &#91;&#93;&#42;sql.ColumnType Capability Report

- Package "gopkg.in/goracle.v2" (goracle)
- Oracle

<table>
	<thead>
		<tr>
			<th>DDL Definition</th><th>ct.Name</th><th>ct.DBTypeName</th><th>ct.Nullable</th><th>ct.DecimalSize</th><th>ct.Length</th><th>ct.ScanType</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td nowrap><code>column_0 CHAR</code></td>
			<td nowrap><code>COLUMN_0</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_1 CHAR (8)</code></td>
			<td nowrap><code>COLUMN_1</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_2 VARCHAR2 (8)</code></td>
			<td nowrap><code>COLUMN_2</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_3 NCHAR</code></td>
			<td nowrap><code>COLUMN_3</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_4 NCHAR (8)</code></td>
			<td nowrap><code>COLUMN_4</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>32</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_5 NVARCHAR2 (8)</code></td>
			<td nowrap><code>COLUMN_5</code></td>
			<td nowrap><code>NVARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>32</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_6 DATE</code></td>
			<td nowrap><code>COLUMN_6</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_7 TIMESTAMP</code></td>
			<td nowrap><code>COLUMN_7</code></td>
			<td nowrap><code>TIMESTAMP</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_8 TIMESTAMP WITH TIME ZONE</code></td>
			<td nowrap><code>COLUMN_8</code></td>
			<td nowrap><code>TIMESTAMP WITH TIMEZONE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_9 INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>COLUMN_9</code></td>
			<td nowrap><code>INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_10 INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>COLUMN_10</code></td>
			<td nowrap><code>INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_11 BLOB</code></td>
			<td nowrap><code>COLUMN_11</code></td>
			<td nowrap><code>BLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_12 CLOB</code></td>
			<td nowrap><code>COLUMN_12</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_13 CLOB</code></td>
			<td nowrap><code>COLUMN_13</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_14 BFILE</code></td>
			<td nowrap><code>COLUMN_14</code></td>
			<td nowrap><code>BFILE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_15 RAW (8)</code></td>
			<td nowrap><code>COLUMN_15</code></td>
			<td nowrap><code>RAW</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_16 NUMBER</code></td>
			<td nowrap><code>COLUMN_16</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(0,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_17 NUMBER (8)</code></td>
			<td nowrap><code>COLUMN_17</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,0)</code></td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_18 NUMBER (8,4)</code></td>
			<td nowrap><code>COLUMN_18</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,4)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_19 FLOAT</code></td>
			<td nowrap><code>COLUMN_19</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(126,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_20 FLOAT (30)</code></td>
			<td nowrap><code>COLUMN_20</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(30,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_21 BINARY_FLOAT</code></td>
			<td nowrap><code>COLUMN_21</code></td>
			<td nowrap><code>FLOAT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float32</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_22 BINARY_DOUBLE</code></td>
			<td nowrap><code>COLUMN_22</code></td>
			<td nowrap><code>DOUBLE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_23 CHAR NOT NULL</code></td>
			<td nowrap><code>COLUMN_23</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_24 CHAR (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_24</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_25 VARCHAR2 (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_25</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_26 NCHAR NOT NULL</code></td>
			<td nowrap><code>COLUMN_26</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>4</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_27 NCHAR (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_27</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>32</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_28 NVARCHAR2 (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_28</code></td>
			<td nowrap><code>NVARCHAR2</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>32</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_29 DATE NOT NULL</code></td>
			<td nowrap><code>COLUMN_29</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_30 TIMESTAMP NOT NULL</code></td>
			<td nowrap><code>COLUMN_30</code></td>
			<td nowrap><code>TIMESTAMP</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_31 TIMESTAMP WITH TIME ZONE NOT NULL</code></td>
			<td nowrap><code>COLUMN_31</code></td>
			<td nowrap><code>TIMESTAMP WITH TIMEZONE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_32 INTERVAL YEAR TO MONTH NOT NULL</code></td>
			<td nowrap><code>COLUMN_32</code></td>
			<td nowrap><code>INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_33 INTERVAL DAY TO SECOND NOT NULL</code></td>
			<td nowrap><code>COLUMN_33</code></td>
			<td nowrap><code>INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_34 BLOB NOT NULL</code></td>
			<td nowrap><code>COLUMN_34</code></td>
			<td nowrap><code>BLOB</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_35 CLOB NOT NULL</code></td>
			<td nowrap><code>COLUMN_35</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_36 CLOB NOT NULL</code></td>
			<td nowrap><code>COLUMN_36</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_37 BFILE NOT NULL</code></td>
			<td nowrap><code>COLUMN_37</code></td>
			<td nowrap><code>BFILE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>math.MaxInt64</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_38 RAW (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_38</code></td>
			<td nowrap><code>RAW</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_39 NUMBER NOT NULL</code></td>
			<td nowrap><code>COLUMN_39</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(0,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_40 NUMBER (8) NOT NULL</code></td>
			<td nowrap><code>COLUMN_40</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(8,0)</code></td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_41 NUMBER (8,4) NOT NULL</code></td>
			<td nowrap><code>COLUMN_41</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(8,4)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_42 FLOAT NOT NULL</code></td>
			<td nowrap><code>COLUMN_42</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(126,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_43 FLOAT (30) NOT NULL</code></td>
			<td nowrap><code>COLUMN_43</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(30,-127)</code></td>
			<td>-</td>
			<td nowrap><code>goracle.Number</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_44 BINARY_FLOAT NOT NULL</code></td>
			<td nowrap><code>COLUMN_44</code></td>
			<td nowrap><code>FLOAT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float32</code></td>
		</tr>
		<tr>
			<td nowrap><code>column_45 BINARY_DOUBLE NOT NULL</code></td>
			<td nowrap><code>COLUMN_45</code></td>
			<td nowrap><code>DOUBLE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
	</tbody>
</table>

Report for [gopkg.in/goracle.v2](https://github.com/go-goracle/goracle) (goracle)<br/>
Test timestamp 2018-02-27T17:21:35Z<br/>
Generated by [drivercaps](https://github.com/jimsmart/drivercaps)

