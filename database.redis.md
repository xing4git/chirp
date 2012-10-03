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
"*" means uid, feed which at uid. "score" is time, "value" is feed id.  

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


## Fans List  
"*" means uid, fans list of uid. "score" is follow time, "value" is uid.

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:fans:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Follows List  
"*" means uid, follow list of uid. "score" is follow time, "value" is uid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:follows:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Feeds List  
"*" means uid, feeds list of uid. "score" is feed create time, "value" is fid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>user:feeds:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Comments List  
"*" means feed id, comments list of fid. "score" is comment create time, "value" is cid.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>feed:comments:*</td>
    <td>sorted set</td>
  </tr>
</table>


## Forwards List  
"*" means feed id, forward list of fid. "score" is forward feed create time, "value" is forward feed id.  

<table>
  <tr>
  	<th>collection</th>
  	<th>type</th>
  </tr>
  <tr>
    <td>feed:forwards:*</td>
    <td>sorted set</td>
  </tr>
</table>