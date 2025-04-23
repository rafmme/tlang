const editor = CodeMirror.fromTextArea(document.getElementById("code"), {
  lineNumbers: true,
  mode: "javascript",
  theme: "monokai",
  tabSize: 2,
});
