package exportAsciinemaJson

import (
	"encoding/json"
	"github.com/stbuehler/go-termrecording/rawrecording"
	"io"
)

func WriteHTML(writer io.Writer, terminalSize rawrecording.TerminalSize, snapshot *Frame, stdoutLocation string, totalTime float64) error {
	stdoutLoc, err := json.Marshal(stdoutLocation)
	if err != nil {
		return err
	}

	if snapshot == nil {
		snapshot = &Frame{
			Lines: []Line{},
		}
	}
	snapshotJson, err := json.Marshal(snapshot.Lines)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(
		`<html>
<head>
  <link rel="stylesheet" type="text/css" href="https://stbuehler.github.io/asciinema-player/css/asciinema-player.css" />
  <link rel="stylesheet" type="text/css" href="https://stbuehler.github.io/asciinema-player/css/themes/tango.css" />
  <link rel="stylesheet" type="text/css" href="https://stbuehler.github.io/asciinema-player/css/themes/solarized-dark.css" />
  <link rel="stylesheet" type="text/css" href="https://stbuehler.github.io/asciinema-player/css/themes/solarized-light.css" />

  <script src="https://stbuehler.github.io/asciinema-player/vendor/react-0.10.0.js"></script>
  <script src="https://stbuehler.github.io/asciinema-player/vendor/JSXTransformer-0.10.0.js"></script>
  <script src="https://stbuehler.github.io/asciinema-player/vendor/jquery-1.10.0.min.js"></script>
  <script src="https://stbuehler.github.io/asciinema-player/vendor/screenfull.js"></script>
  <script src="https://stbuehler.github.io/asciinema-player/js/asciinema-player.min.js"></script>
</head>
<body>
<div id="player-container"></div>
<script>

  React.renderComponent(
    asciinema.Player({
      autoPlay: false,
      movie: new asciinema.Movie(
        ` + Stringify(terminalSize.Columns) + ", " + Stringify(terminalSize.Rows) + `,
        new asciinema.HttpArraySource(` + string(stdoutLoc) + `, 1),
        ` + string(snapshotJson) + `,
        ` + Stringify(totalTime) + `
      )
    }),
    document.getElementById('player-container')
  );

</script>

</body>
</html>`))
	return err
}
