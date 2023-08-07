import React, { Component } from "react";
import "./style.css"

class ChatInput extends Component {
  render() {
    return (
      // <div className="ChatInput">
      //   <input onKeyDown={this.props.send} />
      // </div>
      <div class="input-group">
      <input type="text" class="form-control" placeholder="Write message..."onKeyDown={this.props.send}/>
      <div class="input-group-append ">
          <span class="input-group-text send-icon "><i class="fa-solid fa-paper-plane"></i>
          </span>
      </div>
  </div>
    );
  }
}

export default ChatInput;