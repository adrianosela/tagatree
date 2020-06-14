import base64 from 'base-64';
import config from '../../Config';

export const login = (email, password) => {
  return new Promise((resolve, reject) => {
    fetch(`${config.API_URL}/login`, {
      method: 'GET',
      headers: {
        'Authorization': 'Basic ' + base64.encode(email + ':' + password)
      },
    })
      .then((response) => {
        if (response.status === 200) {
          resolve(response.json());
        }
        if (response.status === 401) {
          reject( { error: 'unauthorized' });
        }
        reject( { error: 'unhandled error' });
      })
      .catch((error) => {
        console.log(error);
        reject({ error: 'Unknown error' });
      });
  });
};
