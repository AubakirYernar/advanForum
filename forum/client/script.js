document.addEventListener("DOMContentLoaded", () => {
    const messageInput = document.getElementById("messageInput");
    const sendButton = document.getElementById("sendButton");
    const messagesDiv = document.getElementById("messages");
  
    const fetchMessages = async () => {
      const response = await fetch("/api/messages");
      const messages = await response.json();
      messagesDiv.innerHTML = messages.map(msg => `<p>${msg.content}</p>`).join("");
    };
  
    const sendMessage = async () => {
      const content = messageInput.value;
      if (!content.trim()) return;
  
      await fetch("/api/messages", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ content }),
      });
  
      messageInput.value = "";
      fetchMessages();
    };
  
    sendButton.addEventListener("click", sendMessage);
    fetchMessages();
  });
  