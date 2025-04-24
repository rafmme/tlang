CodeMirror.defineMode("tlang", function (config, parserConfig) {
  return {
    token: function (stream, state) {
      if (
        stream.match(
          /\b(if|else|fxn|true|false|return|display|make|push|count|first|last|rest|null)\b/
        )
      ) {
        return "keyword";
      }

      if (stream.match(/"[^"]*"/)) {
        return "string";
      }

      if (stream.match(/'[^']*'/)) {
        return "string";
      }

      if (stream.match(/`[^`]*`/)) {
        return "string";
      }

      stream.next();
      return null;
    },
  };
});

CodeMirror.defineMIME("text/tlang", "tlang");

const editor = CodeMirror.fromTextArea(document.getElementById("code"), {
  lineNumbers: true,
  mode: "text/tlang",
  theme: "monokai",
  tabSize: 2,
});
