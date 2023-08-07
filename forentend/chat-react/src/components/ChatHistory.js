import React, { Component } from "react";
import Message from "./message";
import "./style.css"

class ChatHistory extends Component {
  render() {
    // console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map((msg)=><Message message={msg.data} />)
    // console.log("message Formate :- ",messages)
    const formattedMessages = messages.map((msg, index) => (
      <div key={index} className={index % 2 === 0 ? "left" : "right"}>
        {msg}
      </div>
    ));
    return <div><p className="message-content">{formattedMessages}</p></div>;                                                     
  }
}

export default ChatHistory;