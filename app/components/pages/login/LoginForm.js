import React from 'react';
import {
  StyleSheet,
  View,
  Text,
  TextInput,
  TouchableOpacity
} from 'react-native';

export default class LoginForm extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <TextInput
	  placeholder='username or email'
	  placeholderTextColor='#000000'
	  returnKeyType='next'
	  onSubmitEditing={() => this.passwordInput.focus()}
	  keyboardType='email-address'
	  autoCapitalize='none'
	  autoCorrect={false}
          style={styles.input}
        />
        <TextInput
	  placeholder='password'
	  placeholderTextColor='#000000'
	  returnKeyType='go'
	  secureTextEntry
	  ref={(input) => this.passwordInput = input }
          style={styles.input}
        />
	<TouchableOpacity style={styles.buttonContainer}>
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
    fontWeight: "700",
    color: '#fffff0'
  }
});
