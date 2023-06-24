import React, { useEffect, useState } from 'react';
import axios from 'axios';
import API_BASE_URL from './config';

function App() {
  const [blogs, setBlogs] = useState([]);

  useEffect(() => {
    axios.get(`${API_BASE_URL}/blogs`)
      .then(response => {
        setBlogs(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  return (
    <div>
      <h1>Blogs:</h1>
      {blogs.map(blog => (
        <div key={blog.ID}>
          <h3>{blog.Title}</h3>
          <p>{blog.Content}</p>
          <p>Author: {blog.Author}</p>
          <p>Updated: {blog.UpdatedAt}</p>

          {/* Add additional blog details as needed */}
        </div>
      ))}
    </div>
  );
}

export default App;
