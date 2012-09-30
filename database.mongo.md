## Feed  
table name: feed  
index: {fid:1}  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>fid</td>
    <td>string</td>
    <td>feed id</td>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>user id</td>
  </tr>
  <tr>
    <td>content</td>
    <td>FeedContent</td>
    <td>include text and image url</td>
  </tr>
  <tr>
    <td>refid</td>
    <td>string</td>
    <td>ref feed id</td>
  </tr>
  <tr>
    <td>ctime</td>
    <td>int64</td>
    <td>create time, millisecond</td>
  </tr>
</table>


## Feed Deleted  
table name: feed_del  
others like Feed  


## Feed Location  
table name: feedLoc  
index: {loc.loc:'2d'}  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>fid</td>
    <td>string</td>
    <td>feed id</td>
  </tr>
  <tr>
    <td>loc</td>
    <td>Location</td>
    <td>{ctime:int64, loc:[lat, lon]}</td>
  </tr>
</table>


## Feed Location Deleted  
table name: feedLoc_del  
others like Feed Location  


## User  
table name: user  
index: {uid:1}, {username:1}(unique:true), {email:1}(unique:true)  
all the fields are required.  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>user id</td>
  </tr>
  <tr>
    <td>username</td>
    <td>string</td>
    <td></td>
  </tr>
  <tr>
    <td>email</td>
    <td>string</td>
    <td></td>
  </tr>
  <tr>
    <td>pwd</td>
    <td>string</td>
    <td>encrypted password</td>
  </tr>
  <tr>
    <td>avatar</td>
    <td>string</td>
    <td>avatar url</td>
  </tr>
  <tr>
    <td>sex</td>
    <td>int</td>
    <td>0 means man, 1 means woman, 2 means others</td>
  </tr>
  <tr>
    <td>rtime</td>
    <td>int64</td>
    <td>register time, millisecond</td>
  </tr>
</table>


## User Deleted  
table name: user_del  
others like User  


## User Expand  
table name: usreExpand  
index: {uid:1}  
all the fields are optional.  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>user id</td>
  </tr>
  <tr>
    <td>blog</td>
    <td>string</td>
    <td>blog address</td>
  </tr>
  <tr>
    <td>address</td>
    <td>string</td>
    <td></td>
  </tr>
  <tr>
    <td>birthday</td>
    <td>string</td>
    <td>yyyy-mm-dd</td>
  </tr>
  <tr>
    <td>phone</td>
    <td>string</td>
    <td>phone number</td>
  </tr>
  <tr>
    <td>qq</td>
    <td>string</td>
    <td>qq account</td>
  </tr>
  <tr>
    <td>msn</td>
    <td>string</td>
    <td>msn account</td>
  </tr>
  <tr>
    <td>description</td>
    <td>string</td>
    <td>self descrition</td>
  </tr>
  <tr>
    <td>logcnt</td>
    <td>int</td>
    <td>login count</td>
  </tr>
  <tr>
    <td>lltime</td>
    <td>int64</td>
    <td>last login time</td>
  </tr>
</table>


## User Expand Deleted  
table name: usreExpand_del  
others like User Expand  


## User Location  
table name: userLoc  
index: {lloc.loc.:'2d'}  
Location: {ctime:int64, loc:[lat, lon]}  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>feed id</td>
  </tr>
  <tr>
    <td>lloc</td>
    <td>Location</td>
    <td>last location</td>
  </tr>
  <tr>
    <td>hloc</td>
    <td>Location array</td>
    <td>history locations</td>
  </tr>
</table>


## User Location Deleted  
table name: userLoc_del  
others like User Location  


## Comment  
table name: comment  
index: {cid:1}  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>cid</td>
    <td>string</td>
    <td>comment id</td>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>user id</td>
  </tr>
  <tr>
    <td>fid</td>
    <td>string</td>
    <td>feed id</td>
  </tr>
  <tr>
    <td>content</td>
    <td>string</td>
    <td>comment content</td>
  </tr>
  <tr>
    <td>ctime</td>
    <td>int64</td>
    <td>create time, millisecond</td>
  </tr>
</table>


## Comment Deleted  
table name: comment_del  
others like Comment  


## Follow  
table name: follow  
index: {uid:1, beuid:1}(unique:true)  

<table>
  <tr>
  	<th>field</th>
  	<th>type</th>
  	<th>descrition</th>
  </tr>
  <tr>
    <td>uid</td>
    <td>string</td>
    <td>uid follows beuid</td>
  </tr>
  <tr>
    <td>beuid</td>
    <td>string</td>
    <td>the user id who was followed</td>
  </tr>
  <tr>
    <td>ctime</td>
    <td>int64</td>
    <td>create time, millisecond</td>
  </tr>
</table>