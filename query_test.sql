SELECT 
  child.id,
  child.name,
  parent.name AS parent_name
FROM 
  your_table_name AS child
LEFT JOIN 
  your_table_name AS parent
ON 
  child.parent_id = parent.id;
