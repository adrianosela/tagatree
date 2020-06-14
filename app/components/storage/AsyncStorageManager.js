import React from 'react';
import AsyncStorage from '@react-native-community/async-storage';

export default class AsyncStorageManager {
    static instance = null;

    static getInstance() {
        if (AsyncStorageManager.instance == null) {
            AsyncStorageManager.instance = new AsyncStorageManager();
        }

        return AsyncStorageManager.instance;
    }

    _storeData = async (key, val) => {
        try {
            await AsyncStorage.setItem(key, val);
        } catch (error) {
            console.log(error);
        }
    };

    _retrieveData = async (key) => {
        try {
            return await AsyncStorage.getItem(key);
        } catch (error) {
            console.log(error);
            return null;
        }
    };

    _removeData = async (key) => {
        try {
            return await AsyncStorage.removeItem(key);
        } catch (error) {
            console.log(error);
        }
    };

    saveUserToken = userToken => {
        return this._storeData("userToken", userToken);
    };

    getUserToken = () => {
        return this._retrieveData("userToken");
    };

    clearUserToken = () => {
        return this._removeData("userToken");  
    }
}
