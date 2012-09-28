## Feed  
table name: feed  
index: {fid:1} 

<table>
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
  <tr>
    <td>status</td>
    <td>int</td>
    <td>0 means ok, 1 means deleted</td>
  </tr>
</table>


## FeedLoc  
table name: feedLoc  
index: {loc:'2d'}  

<table>
  <tr>
    <td>fid</td>
    <td>string</td>
    <td>feed id</td>
  </tr>
  <tr>
    <td>loc</td>
    <td>array</td>
    <td>[lat, lon]</td>
  </tr>
  <tr>
    <td>status</td>
    <td>int</td>
    <td>0 means ok, 1 means deleted</td>
  </tr>
</table>


## User  
table name: user  
index: {uid:1}, {username:1}, {email:1}  

<table>
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
    <td>rtime</td>
    <td>int64</td>
    <td>register time, millisecond</td>
  </tr>
  <tr>
    <td>status</td>
    <td>int</td>
    <td>0 means ok, 1 means deleted</td>
  </tr>
</table>


## Comment  
table name: comment  
index: {cid:1}  

<table>
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
  <tr>
    <td>status</td>
    <td>int</td>
    <td>0 means ok, 1 means deleted</td>
  </tr>
</table>


## Follow  
table name: follow  
index: {uid:1, beuid:1}(unique:true)  

<table>
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