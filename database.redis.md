## Time Line  
"*" means uid, homeline of uid. "score" is feed create time, "value" is feed id.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:timeline:*</td>
    <td>sorted set</td>
  </tr>
</table>


## At  
"*" means uid, feeds or comments which @ uid. "score" is @ time, "value" contains @ type(feed or comment) and id(feed id or comment id).  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:at:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Fan List  
"*" means uid, fans list of uid. "score" is follow time, "value" is uid.

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:fan:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Follow List  
"*" means uid, follow list of uid. "score" is follow time, "value" is uid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:follow:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Feed List  
"*" means uid, feeds list of uid. "score" is feed create time, "value" is fid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:feed:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Comment List  
"*" means feed id, comments list of fid. "score" is comment create time, "value" is cid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>feed:comment:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Forward List  
"*" means feed id, forward list of fid. "score" is forward feed create time, "value" is forward feed id.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>feed:forward:*</td>
    <td>sorted set</td>
  </tr>
</table>