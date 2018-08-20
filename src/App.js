import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { ApolloProvider } from 'react-apollo';
import client from './apollo';
import Poll from './Poll';
import { Users } from './Users.react'
import { getUserId } from './session'

class App extends Component {
  constructor (props) {
    super(props);
    this.state = {loading: true, userId: ''};
  }

  componentWillMount() {
    getUserId().then((userId) => {
      this.setState({loading: false, userId});
    });
  }

  render() {
    if (this.state.loading) return <p>Loading...</p>;
    return (
      <ApolloProvider client={client}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h1 className="App-title">Welcome to React</h1>
          </header>
          <Poll userId={this.state.userId}/>
          <Users />
        </div>
      </ApolloProvider>
    );
  }
}

export default App;
