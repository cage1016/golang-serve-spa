import React from 'react';

const NotFound = () => (
  <div className="notfound">
    <h3>404 page not found</h3>
    <p>We are sorry but the page you are looking for does not exist.</p>
  </div>
)

// export default NotFound;
export default {
  path : '*',
  component : NotFound
}
