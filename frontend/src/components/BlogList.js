import React, { useEffect, useState } from 'react';
import axios from 'axios';
import API_BASE_URL from '../config.js';

function BlogList() {
  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/blogs`);
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return <div>Blog List</div>;
}

export default BlogList;
