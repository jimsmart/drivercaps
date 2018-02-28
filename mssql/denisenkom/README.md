
# Driver sql.ColumnType Capability Report

- Package "github.com/denisenkom/go-mssqldb" (mssql)
- Microsoft SQL Server

<table>
	<thead>
		<tr>
			<th>DDL Definition</th><th>.Name</th><th>.DBTypeName</th><th>.Nullable</th><th>.DecimalSize</th><th>.Length</th><th>.ScanType</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td nowrap><code>t_0 bigint</code></td>
			<td nowrap><code>t_0</code></td>
			<td nowrap><code>BIGINT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_1 bit</code></td>
			<td nowrap><code>t_1</code></td>
			<td nowrap><code>BIT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>bool</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_2 decimal</code></td>
			<td nowrap><code>t_2</code></td>
			<td nowrap><code>DECIMAL</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(18,0)</code></td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_3 int</code></td>
			<td nowrap><code>t_3</code></td>
			<td nowrap><code>INT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_4 numeric</code></td>
			<td nowrap><code>t_4</code></td>
			<td nowrap><code>DECIMAL</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(18,0)</code></td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_5 smallint</code></td>
			<td nowrap><code>t_5</code></td>
			<td nowrap><code>SMALLINT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_6 tinyint</code></td>
			<td nowrap><code>t_6</code></td>
			<td nowrap><code>TINYINT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_7 float</code></td>
			<td nowrap><code>t_7</code></td>
			<td nowrap><code>FLOAT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_8 real</code></td>
			<td nowrap><code>t_8</code></td>
			<td nowrap><code>REAL</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_9 date</code></td>
			<td nowrap><code>t_9</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_10 datetime2</code></td>
			<td nowrap><code>t_10</code></td>
			<td nowrap><code>DATETIME2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_11 datetime</code></td>
			<td nowrap><code>t_11</code></td>
			<td nowrap><code>DATETIME</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_12 datetimeoffset</code></td>
			<td nowrap><code>t_12</code></td>
			<td nowrap><code>DATETIMEOFFSET</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_13 smalldatetime</code></td>
			<td nowrap><code>t_13</code></td>
			<td nowrap><code>SMALLDATETIME</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_14 time</code></td>
			<td nowrap><code>t_14</code></td>
			<td nowrap><code>TIME</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_15 char</code></td>
			<td nowrap><code>t_15</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_16 text</code></td>
			<td nowrap><code>t_16</code></td>
			<td nowrap><code>TEXT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>2147483647</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_17 varchar</code></td>
			<td nowrap><code>t_17</code></td>
			<td nowrap><code>VARCHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_18 nchar</code></td>
			<td nowrap><code>t_18</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_19 ntext</code></td>
			<td nowrap><code>t_19</code></td>
			<td nowrap><code>NTEXT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1073741823</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_20 nvarchar</code></td>
			<td nowrap><code>t_20</code></td>
			<td nowrap><code>NVARCHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_21 binary</code></td>
			<td nowrap><code>t_21</code></td>
			<td nowrap><code>BINARY</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_22 image</code></td>
			<td nowrap><code>t_22</code></td>
			<td nowrap><code>IMAGE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>2147483647</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_23 varbinary</code></td>
			<td nowrap><code>t_23</code></td>
			<td nowrap><code>VARBINARY</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_24 sql_variant</code></td>
			<td nowrap><code>t_24</code></td>
			<td nowrap><code>SQL_VARIANT</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>nil</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_25 rowversion</code></td>
			<td nowrap><code>t_25</code></td>
			<td nowrap><code>BINARY</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_26 uniqueidentifier</code></td>
			<td nowrap><code>t_26</code></td>
			<td nowrap><code>UNIQUEIDENTIFIER</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_27 xml</code></td>
			<td nowrap><code>t_27</code></td>
			<td nowrap><code>XML</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1073741822</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_28 bigint NOT NULL</code></td>
			<td nowrap><code>t_28</code></td>
			<td nowrap><code>BIGINT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_29 bit NOT NULL</code></td>
			<td nowrap><code>t_29</code></td>
			<td nowrap><code>BIT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>bool</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_30 decimal NOT NULL</code></td>
			<td nowrap><code>t_30</code></td>
			<td nowrap><code>DECIMAL</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(18,0)</code></td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_31 int NOT NULL</code></td>
			<td nowrap><code>t_31</code></td>
			<td nowrap><code>INT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_32 numeric NOT NULL</code></td>
			<td nowrap><code>t_32</code></td>
			<td nowrap><code>DECIMAL</code></td>
			<td nowrap><code>false</code></td>
			<td nowrap><code>(18,0)</code></td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_33 smallint NOT NULL</code></td>
			<td nowrap><code>t_33</code></td>
			<td nowrap><code>SMALLINT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_34 tinyint NOT NULL</code></td>
			<td nowrap><code>t_34</code></td>
			<td nowrap><code>TINYINT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>int64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_35 float NOT NULL</code></td>
			<td nowrap><code>t_35</code></td>
			<td nowrap><code>FLOAT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_36 real NOT NULL</code></td>
			<td nowrap><code>t_36</code></td>
			<td nowrap><code>REAL</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_37 date NOT NULL</code></td>
			<td nowrap><code>t_37</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_38 datetime2 NOT NULL</code></td>
			<td nowrap><code>t_38</code></td>
			<td nowrap><code>DATETIME2</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_39 datetime NOT NULL</code></td>
			<td nowrap><code>t_39</code></td>
			<td nowrap><code>DATETIME</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_40 datetimeoffset NOT NULL</code></td>
			<td nowrap><code>t_40</code></td>
			<td nowrap><code>DATETIMEOFFSET</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_41 smalldatetime NOT NULL</code></td>
			<td nowrap><code>t_41</code></td>
			<td nowrap><code>SMALLDATETIME</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_42 time NOT NULL</code></td>
			<td nowrap><code>t_42</code></td>
			<td nowrap><code>TIME</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_43 char NOT NULL</code></td>
			<td nowrap><code>t_43</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_44 text NOT NULL</code></td>
			<td nowrap><code>t_44</code></td>
			<td nowrap><code>TEXT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>2147483647</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_45 varchar NOT NULL</code></td>
			<td nowrap><code>t_45</code></td>
			<td nowrap><code>VARCHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_46 nchar NOT NULL</code></td>
			<td nowrap><code>t_46</code></td>
			<td nowrap><code>NCHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_47 ntext NOT NULL</code></td>
			<td nowrap><code>t_47</code></td>
			<td nowrap><code>NTEXT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1073741823</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_48 nvarchar NOT NULL</code></td>
			<td nowrap><code>t_48</code></td>
			<td nowrap><code>NVARCHAR</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_49 binary NOT NULL</code></td>
			<td nowrap><code>t_49</code></td>
			<td nowrap><code>BINARY</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_50 image NOT NULL</code></td>
			<td nowrap><code>t_50</code></td>
			<td nowrap><code>IMAGE</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>2147483647</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_51 varbinary NOT NULL</code></td>
			<td nowrap><code>t_51</code></td>
			<td nowrap><code>VARBINARY</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_52 sql_variant NOT NULL</code></td>
			<td nowrap><code>t_52</code></td>
			<td nowrap><code>SQL_VARIANT</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>nil</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_53 uniqueidentifier NOT NULL</code></td>
			<td nowrap><code>t_53</code></td>
			<td nowrap><code>UNIQUEIDENTIFIER</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td>-</td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_54 xml NOT NULL</code></td>
			<td nowrap><code>t_54</code></td>
			<td nowrap><code>XML</code></td>
			<td nowrap><code>false</code></td>
			<td>-</td>
			<td nowrap><code>1073741822</code></td>
			<td nowrap><code>string</code></td>
		</tr>
	</tbody>
</table>

Report for [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) (mssql)<br/>
Test timestamp 2018-02-28T19:53:57Z<br/>
Generated by [drivercaps](https://github.com/jimsmart/drivercaps)

