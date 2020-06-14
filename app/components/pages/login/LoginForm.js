import React from 'react';
import {
  StyleSheet,
  View,
  Text,
  TextInput,
  TouchableOpacity
} from 'react-native';
import * as auth from '../../api/Auth';
import AsyncStorageManager from '../../storage/AsyncStorageManager';

export default class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.onLoginButtonPressed = this.onLoginButtonPressed.bind(this);
    this.state = {email: '', password: ''};
  }

  onLoginButtonPressed() {
    if (!this.state.email || !this.state.password) {
      // fixme: use pretty alerts/errors
      console.log('no email or password!');
      return;
    }
    auth.login(this.state.email, this.state.password)
      .then((response) => {
        AsyncStorageManager.getInstance().saveUserToken(response.token);
        this.props.onSuccess();
      })
      .catch((ex) => {
        console.log(ex);
      });
  }

  render() {
    return (
      <View style={styles.container}>
        <TextInput
          placeholder='username or email'
          onChangeText={(val) => this.setState({email: val})}
          placeholderTextColor='#000000'
          returnKeyType='next'
          keyboardType='email-address'
          autoCapitalize='none'
          autoCorrect={false}
          onSubmitEditing={() => this.passwordInput.focus()}
          style={styles.input}
        />
        <TextInput
          placeholder='password'
          onChangeText={(val) => this.setState({password: val})}
          placeholderTextColor='#000000'
          returnKeyType='go'
          secureTextEntry
          ref={(input) => this.passwordInput = input }
          style={styles.input}
        />
        <TouchableOpacity style={styles.buttonContainer} onPress={this.onLoginButtonPressed}>
          <Text style={styles.buttonText}>Login</Text>
        </TouchableOpacity>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    padding: 20,
    paddingHorizontal: 50
  },
  input: {
    height: 40,
    marginBottom: 20,
    backgroundColor: '#00b894',
    color: '#000000',
    paddingHorizontal: 20,
    textAlign: 'center'
  },
  buttonContainer: {
    backgroundColor: '#00cec9',
    paddingVertical: 10
  },
  buttonText: {
    textAlign: 'center',
    fontSize: 20,
    fontWeight: '700',
    color: '#fffff0'
  }
});
