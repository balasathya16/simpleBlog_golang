import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Blog = () => {
  const [blogs, setBlogs] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/blogs')
      .then(response => {
        setBlogs(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  return (
    <div>
      <h1>Blog Page</h1>
      {blogs.map(blog => (
        <div key={blog.ID}>
          <h3>{blog.Title}</h3>
          <p>{blog.Content}</p>
          <p>Author: {blog.Author}</p>
          <p>Updated: {blog.UpdatedAt}</p>
        </div>
      ))}
    </div>
  );
};

export default Blog;
