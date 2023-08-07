import logo from './logo.svg';
import './App.css';
import "../src/components/style.css"
import { connect, sendMsg } from './api';
import { Component } from 'react';
import { Header } from './components/header';
import ChatHistory from './components/ChatHistory'
import ChatInput from './components/ChatInput';
// import socketIOClient, { Socket } from "socket.io-client"
// const socket=socketIOClient("http://localhost:3000/")
// console.log("welcome",socket)
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    }
  }
  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }
  componentDidMount() {
    // socket.on("new message",data=>{
    //   console.log("data:-",data)
    // })
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  render() {
    return (
      <>
        <div className='App'>
          <div class="container">
            <div class="chat-page">
              <div class="msg-inbox">
                <div class="chats">
                  <div class="msg-page">
                    <div class="received-chats">
                      <div class="received-msg">
                        <div class="received-msg-inbox">
                          <ChatHistory chatHistory={this.state.chatHistory} />
                        </div>
                      </div>
                    </div>

                  </div>
                </div>


                <div class="msg-bottom">
                  <ChatInput send={this.send}/>

                </div>

              </div>
            </div>
          </div>
        </div>
      </>
    );
  }
}
export default App;
{/* <div className='middle'>
            <div className='main'>
              <Header />
              <ChatHistory chatHistory={this.state.chatHistory} />
            </div>
            <div className='chat'>
              <ChatInput send={this.send} />
            </div>
          </div> */}
{/* <button onClick={this.send}>Click</button> */ }