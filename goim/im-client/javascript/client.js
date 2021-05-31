(function(win) {
    const rawHeaderLen = 16;
    const packetOffset = 0;
    const headerOffset = 4;
    const verOffset = 6;
    const opOffset = 8;
    const seqOffset = 12;

    let Client = function(options) {
        const MAX_CONNECT_TIMES = 10;
        const DELAY = 15000;
        this.options = options || {};
        this.createConnect(MAX_CONNECT_TIMES, DELAY);
    }

    let appendMsg = function(text) {
        let span = document.createElement("SPAN");
        var text = document.createTextNode(text);
        span.appendChild(text);
        document.getElementById("box").appendChild(span);
    }

    Client.prototype.createConnect = function(max, delay) {
        let self = this;
        if (max === 0) {
            return;
        }
        connect();

        let textDecoder = new TextDecoder();
        let textEncoder = new TextEncoder();
        let heartbeatInterval;
        function connect() {
            let ws = new WebSocket('ws://127.0.0.1:3102/sub');
            //var ws = new WebSocket('ws://127.0.0.1:3102/sub');
            ws.binaryType = 'arraybuffer';
            ws.onopen = function() {
                auth();
            }

            ws.onmessage = function(evt) {
                let data = evt.data;
                let dataView = new DataView(data, 0);
                let packetLen = dataView.getInt32(packetOffset);
                let headerLen = dataView.getInt16(headerOffset);
                let ver = dataView.getInt16(verOffset);
                let op = dataView.getInt32(opOffset);
                let seq = dataView.getInt32(seqOffset);

                console.log("receiveHeader: packetLen=" + packetLen, "headerLen=" + headerLen, "ver=" + ver, "op=" + op, "seq=" + seq);

                switch(op) {
                    case 8:
                        // auth reply ok
                        document.getElementById("status").innerHTML = "<color style='color:green'>ok<color>";
                        appendMsg("receive: auth reply");
                        // send a heartbeat to server
                        heartbeat();
                        heartbeatInterval = setInterval(heartbeat, 30 * 1000);
                        break;
                    case 3:
                        // receive a heartbeat from server
                        console.log("receive: heartbeat");
                        appendMsg("receive: heartbeat reply");
                        break;
                    case 9:
                        // batch message
                        for (let offset=rawHeaderLen; offset<data.byteLength; offset+=packetLen) {
                            // parse
                            let packetLen = dataView.getInt32(offset);
                            let headerLen = dataView.getInt16(offset+headerOffset);
                            let ver = dataView.getInt16(offset+verOffset);
                            let op = dataView.getInt32(offset+opOffset);
                            let seq = dataView.getInt32(offset+seqOffset);
                            let msgBody = textDecoder.decode(data.slice(offset+headerLen, offset+packetLen));
                            // callback
                            messageReceived(ver, msgBody);
                            appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        }
                        break;
                    default:
                        let msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
                        messageReceived(ver, msgBody);
                        appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        break
                }
            }

            ws.onclose = function() {
                if (heartbeatInterval) clearInterval(heartbeatInterval);
                setTimeout(reConnect, delay);

                document.getElementById("status").innerHTML =  "<color style='color:red'>failed<color>";
            }

            function heartbeat() {
                let headerBuf = new ArrayBuffer(rawHeaderLen);
                let headerView = new DataView(headerBuf, 0);
                headerView.setInt32(packetOffset, rawHeaderLen);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 2);
                headerView.setInt32(seqOffset, 1);
                ws.send(headerBuf);
                console.log("send: heartbeat");
                appendMsg("send: heartbeat");
            }

            function auth() {
                let token = '{"mid":123, "room_id":"live://1000", "platform":"web", "accepts":[1000,1001,1002]}'
                let headerBuf = new ArrayBuffer(rawHeaderLen);
                let headerView = new DataView(headerBuf, 0);
                let bodyBuf = textEncoder.encode(token);
                headerView.setInt32(packetOffset, rawHeaderLen + bodyBuf.byteLength);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 7);
                headerView.setInt32(seqOffset, 1);
                ws.send(mergeArrayBuffer(headerBuf, bodyBuf));
                debugger
                appendMsg("send: auth token: " + token);
            }

            function messageReceived(ver, body) {
                let notify = self.options.notify;
                if(notify) notify(body);
                console.log("messageReceived:", "ver=" + ver, "body=" + body);
            }

            function mergeArrayBuffer(ab1, ab2) {
                let u81 = new Uint8Array(ab1),
                    u82 = new Uint8Array(ab2),
                    res = new Uint8Array(ab1.byteLength + ab2.byteLength);
                res.set(u81, 0);
                res.set(u82, ab1.byteLength);
                console.log(res)
                return res.buffer;
            }

            function char2ab(str) {
                let buf = new ArrayBuffer(str.length);
                let bufView = new Uint8Array(buf);
                for (let i=0; i<str.length; i++) {
                    bufView[i] = str[i];
                }
                return buf;
            }

        }

        function reConnect() {
            self.createConnect(--max, delay * 2);
        }
    }

    win['MyClient'] = Client;
})(window);
