package main

const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Eric's todo</title>
        <style> 
        input[type=checkbox]:checked + input[type=text].strikethrough {
            text-decoration: line-through;
            color: var(--bs-gray);
            text-decoration-color: var(--bs-gray);
        }
        </style>
    </head>
    <body>
        {{range $i, $a := .}}
            <div>
                <input id="formCheckBox-1" type="checkbox" />
                <input class="strikethrough" type="text" value="{{.Description}}" for="formCheckBox-1" />
            </div>
        {{else}}
            <div><strong>no rows, click the + button</strong></div>
        {{end}}
        <svg class="bi bi-plus-circle justify-content-md-end align-items-md-end" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" viewBox="0 0 16 16" style="font-size: 55px;position: absolute;bottom: 25px;right: 25px;">
            <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"></path>
            <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z"></path>
        </svg>
    </body>
</html>`