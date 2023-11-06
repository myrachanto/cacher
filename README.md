# cacher
## package to cache authorization module rights

while building a big system you need a means to cache authorization modules otherwise the options
are 

1. Embeb every request with ginormous amount of unnecesary info in a token . for example like category module user has read and write rights and
   every other module the system implements

   or
   
2. With every request you receive from a the from end you query the database for the access rights of the module the user is looking to get or update or create 
  something. which will amount to a lot of unnecesary traffic to the database

and thats why i created this package. 

Its simple it implements a cache that store the access rights.

it has three methods 
1. Put
2. Get
3. Invalidate

It utilizes the second option of quering the database but regulates the amount of database query being sent. the database query is only triggered when
the user in question is not cached
