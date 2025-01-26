# Practise API Server with Dummy Data and DB in Golang


A simple CRUD API Server in GoLang which will contain endpoints for creating, reading, updating and deleting products along with their categories and manufacturers.

## Steps
1. Create a new project
2. Create a new Go file
3. Add the necessary imports
4. Create a new router
5. Create a new DB (dummy)
6. Create controllers for CRUD operations
7. Run the server



**Theme** : I will be creating a simple online store API which will contain products, manufacturer of those products and category of those products




## Task List Example

- [x] Create a new project and then a new Go file and initially keep everything in a single file
- [x] Add the necessary imports.
- [x] Create structs for the necessary fields.
- [x] Create a new DB (dummy).
- [ ] Create POST endpoint **{/api/products}** and its respective controller for creating/adding a new product from manufacturer side.
- [x] Create GET endpoint **{/api/products}** and its respective controller for fetching all the products from the DB.
- [x] Create a GET endpoint **{/api/products/{id}}** and its respective controller for fetching a product by its id.
- [ ] Create a PUT endpoint **{/api/products/{id}}** and its respective controller for updating a product by its id.
- [ ] Create a DELETE endpoint **{/api/products/{id}}** and its respective controller for deleting a product by its id.
- [ ] Now run the server and test the endpoints.
- [ ] Now start to break down the controllers and routes into seperate files and folders.
- [ ] Replace the dummy DB with actual DB.
- [ ] Add a new endpoint **{/api/products/{id}/category}** and its respective controller for fetching a product by its id and its category.
- [ ] Add input validation middleware for POST and PUT requests
- [ ] Implement proper error handling and custom error responses

Note: Use `- [x]` for completed items and `- [ ]` for incomplete items



