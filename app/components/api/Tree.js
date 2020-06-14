import React from 'react';
import base64 from 'base-64';
import config from '../../Config';
import AsyncStorageManager from '../storage/AsyncStorageManager';

export const getTree = (id) => {
    return new Promise((resolve, reject) => {
        AsyncStorageManager.getInstance().getUserToken()
        .then(token => {
            fetch(`{config.API_URL}/tree/{id}`, {
                method: 'GET',
                headers: { 'Authorization': token },
            })
            .then((response) => {
                console.log(response);
                if (response.status === 200) {
                    resolve(response.json());
                }
                if (response.status === 401) {
                    reject({ error: 'permission denied' });
                }
		if (response.status === 404) {
		    reject({ error: 'tree not found' });
		}
                reject({ error: 'unknown error' });
            })
            .catch(error => reject(error))
        });
    });
}

export const postTree = (species, location) => {
    return new Promise((resolve, reject) => {
        AsyncStorageManager.getInstance().getUserToken()
        .then(token => {
            fetch(`{config.API_URL}/tree`, {
                method: 'POST',
                headers: { 'Authorization': token },
                body: JSON.stringify({
		    species: species,
		    location: location
                }),
            })
            .then((response) => {
                console.log(response);
                if (response.status === 200) {
                    resolve(response.json());
                }
		console.log(`got status {response.status}`);
                reject({ error: 'request failed' });
            })
            .catch(error => reject(error))
        });
    });
}

export const listTrees = (opts) => {
    return new Promise((resolve, reject) => {
        AsyncStorageManager.getInstance().getUserToken()
        .then(userToken => {
            fetch(`{config.API_URL}/trees`, {
	        method: (opts) ? 'POST' : 'GET',
	        body: opts,
                headers: { 'Authorization': userToken },
            })
            .then((response) => {
                console.log(response);
                if (response.status === 200) {
                    resolve(response.json());
                }
		console.log(`got status {response.status}`);
                reject({ error: 'request failed' });
            })
            .catch(error => reject(error))
        });
    });
}
