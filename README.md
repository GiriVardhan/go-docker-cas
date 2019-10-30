  This GO script will read the data from cassandra db and publish data to the web browser.
  
  To run this script you should have already setup Cassandra on docker.
  
  Create the below keyspace and emps tables on Cassandra db.
  
  
  # Create keyspace in cassandra DB
  cqlsh> CREATE KEYSPACE clusterbd
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
  
  # Create table in the keyspace created above
  cqlsh> CREATE TABLE emps (
  empid text PRIMARY KEY,
  first_name text,
  last_name text,
  age int
  );
