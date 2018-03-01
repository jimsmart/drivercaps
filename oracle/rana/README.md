
# Driver sql.ColumnType Capability Report

- Package "gopkg.in/rana/ora.v4" (ora)
- Oracle

<table>
	<thead>
		<tr>
			<th>DDL Definition</th><th>.Name</th><th>.DBTypeName</th><th>.Nullable</th><th>.DecimalSize</th><th>.Length</th><th>.ScanType</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td nowrap><code>t_0 CHAR</code></td>
			<td nowrap><code>T_0</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_1 CHAR (8)</code></td>
			<td nowrap><code>T_1</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_2 VARCHAR2 (8)</code></td>
			<td nowrap><code>T_2</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_3 NCHAR</code></td>
			<td nowrap><code>T_3</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>2</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_4 NCHAR (8)</code></td>
			<td nowrap><code>T_4</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>16</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_5 NVARCHAR2 (8)</code></td>
			<td nowrap><code>T_5</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>16</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_6 DATE</code></td>
			<td nowrap><code>T_6</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>7</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_7 TIMESTAMP</code></td>
			<td nowrap><code>T_7</code></td>
			<td nowrap><code>TIMESTAMP</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>11</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_8 TIMESTAMP WITH TIME ZONE</code></td>
			<td nowrap><code>T_8</code></td>
			<td nowrap><code>TIMESTAMP WITH TIME ZONE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>13</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_9 INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>T_9</code></td>
			<td nowrap><code>INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>5</code></td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_10 INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>T_10</code></td>
			<td nowrap><code>INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>11</code></td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_11 BLOB</code></td>
			<td nowrap><code>T_11</code></td>
			<td nowrap><code>BLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_12 CLOB</code></td>
			<td nowrap><code>T_12</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_13 CLOB</code></td>
			<td nowrap><code>T_13</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_14 BFILE</code></td>
			<td nowrap><code>T_14</code></td>
			<td nowrap><code>FILE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>530</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_15 RAW (8)</code></td>
			<td nowrap><code>T_15</code></td>
			<td nowrap><code>RAW</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_16 NUMBER</code></td>
			<td nowrap><code>T_16</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(0,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_17 NUMBER (8)</code></td>
			<td nowrap><code>T_17</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,0)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_18 NUMBER (8,4)</code></td>
			<td nowrap><code>T_18</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,4)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_19 FLOAT</code></td>
			<td nowrap><code>T_19</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(126,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_20 FLOAT (30)</code></td>
			<td nowrap><code>T_20</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(30,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_21 BINARY_FLOAT</code></td>
			<td nowrap><code>T_21</code></td>
			<td nowrap><code>100</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4</code></td>
			<td nowrap><code>nil</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_22 BINARY_DOUBLE</code></td>
			<td nowrap><code>T_22</code></td>
			<td nowrap><code>101</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>nil</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_23 CHAR NOT NULL</code></td>
			<td nowrap><code>T_23</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>1</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_24 CHAR (8) NOT NULL</code></td>
			<td nowrap><code>T_24</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_25 VARCHAR2 (8) NOT NULL</code></td>
			<td nowrap><code>T_25</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_26 NCHAR NOT NULL</code></td>
			<td nowrap><code>T_26</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>2</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_27 NCHAR (8) NOT NULL</code></td>
			<td nowrap><code>T_27</code></td>
			<td nowrap><code>CHAR</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>16</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_28 NVARCHAR2 (8) NOT NULL</code></td>
			<td nowrap><code>T_28</code></td>
			<td nowrap><code>VARCHAR2</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>16</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_29 DATE NOT NULL</code></td>
			<td nowrap><code>T_29</code></td>
			<td nowrap><code>DATE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>7</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_30 TIMESTAMP NOT NULL</code></td>
			<td nowrap><code>T_30</code></td>
			<td nowrap><code>TIMESTAMP</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>11</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_31 TIMESTAMP WITH TIME ZONE NOT NULL</code></td>
			<td nowrap><code>T_31</code></td>
			<td nowrap><code>TIMESTAMP WITH TIME ZONE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>13</code></td>
			<td nowrap><code>time.Time</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_32 INTERVAL YEAR TO MONTH NOT NULL</code></td>
			<td nowrap><code>T_32</code></td>
			<td nowrap><code>INTERVAL YEAR TO MONTH</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>5</code></td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_33 INTERVAL DAY TO SECOND NOT NULL</code></td>
			<td nowrap><code>T_33</code></td>
			<td nowrap><code>INTERVAL DAY TO SECOND</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>11</code></td>
			<td nowrap><code>time.Duration</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_34 BLOB NOT NULL</code></td>
			<td nowrap><code>T_34</code></td>
			<td nowrap><code>BLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_35 CLOB NOT NULL</code></td>
			<td nowrap><code>T_35</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_36 CLOB NOT NULL</code></td>
			<td nowrap><code>T_36</code></td>
			<td nowrap><code>CLOB</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4000</code></td>
			<td nowrap><code>string</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_37 BFILE NOT NULL</code></td>
			<td nowrap><code>T_37</code></td>
			<td nowrap><code>FILE</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>530</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_38 RAW (8) NOT NULL</code></td>
			<td nowrap><code>T_38</code></td>
			<td nowrap><code>RAW</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>[]uint8</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_39 NUMBER NOT NULL</code></td>
			<td nowrap><code>T_39</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(0,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_40 NUMBER (8) NOT NULL</code></td>
			<td nowrap><code>T_40</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,0)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_41 NUMBER (8,4) NOT NULL</code></td>
			<td nowrap><code>T_41</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(8,4)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_42 FLOAT NOT NULL</code></td>
			<td nowrap><code>T_42</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(126,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_43 FLOAT (30) NOT NULL</code></td>
			<td nowrap><code>T_43</code></td>
			<td nowrap><code>NUMBER</code></td>
			<td nowrap><code>true</code></td>
			<td nowrap><code>(30,-127)</code></td>
			<td nowrap><code>22</code></td>
			<td nowrap><code>float64</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_44 BINARY_FLOAT NOT NULL</code></td>
			<td nowrap><code>T_44</code></td>
			<td nowrap><code>100</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>4</code></td>
			<td nowrap><code>nil</code></td>
		</tr>
		<tr>
			<td nowrap><code>t_45 BINARY_DOUBLE NOT NULL</code></td>
			<td nowrap><code>T_45</code></td>
			<td nowrap><code>101</code></td>
			<td nowrap><code>true</code></td>
			<td>-</td>
			<td nowrap><code>8</code></td>
			<td nowrap><code>nil</code></td>
		</tr>
	</tbody>
</table>

Report for [gopkg.in/rana/ora.v4](https://github.com/rana/ora) (ora)<br/>
Test timestamp 2018-03-01T00:12:11Z<br/>
Generated by [drivercaps](https://github.com/jimsmart/drivercaps)

