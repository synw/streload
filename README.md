# Static reload

Watch static files and reload in the browser when a change occurs.

## Install

Use go get and build or grab the binary release

## Usage
   
Add this javascript to the html file you want to reload:

   ```javascript
<script type="text/javascript">
    (function() {
        var conn = new WebSocket("ws://localhost:8042/ws");
        conn.onclose = function(evt) {
            console.log('Connection closed');
        }
        conn.onmessage = function(evt) {
            window.location.reload();
        }
    })();
</script>
   ```

Run the watcher:

   ```
   ./streload path/to/a/file/or/folder
   ```
   
Open the file normaly, it will be reloaded when a change occurs