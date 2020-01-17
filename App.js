import React from 'react';
import {WebView} from 'react-native-webview';

const App: () => React$Node = () => (
  <WebView source={{uri: 'http://localhost:8080/posts'}} />
);

export default App;
