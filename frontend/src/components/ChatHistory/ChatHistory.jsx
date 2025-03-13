import React, { Component } from 'react';
import './ChatHistory.scss';
import Message from '../Message/Message';

class ChatHistory extends Component {
    render() {
        console.log(this.props.chatHistory);
        const messages = this.props.chatHistory.map((msg, key) => <Message message={msg.data} key={key} />);

        return (
            <div className="ChatHistory">
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;