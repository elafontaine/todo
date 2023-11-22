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
        <script>
        $('#modal-add-task').on('shown.bs.modal', function () {
            $('#addIcon').trigger('focus')
          })
        </script>
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
        <svg id="addIcon" class="bi bi-plus-circle justify-content-md-end align-items-md-end" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" viewBox="0 0 16 16" style="font-size: 55px;position: absolute;bottom: 25px;right: 25px;">
            <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"></path>
            <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z"></path>
        </svg>
        <div id="modal-add-task" class="modal fade show" role="dialog" tabindex="-1" style="display: block;">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <form id="task-form" method="POST" action="/add">
                        <div class="modal-header">
                            <h4 class="modal-title">Add Task</h4><button class="btn-close" type="button" aria-label="Close" data-bs-dismiss="modal"></button>
                        </div>
                        <div class="modal-body">
                            <div class="d-sm-flex align-items-sm-end md-form mb-5 input-group-text align-bottom"><label class="form-label" data-error="wrong" date-success="right" for="modalTaskField">Task :</label><input id="modalTaskField" class="form-control form-control validate" type="text" value="task" /></div>
                        </div>
                        <div class="modal-footer"><button class="btn btn-light" type="button" data-bs-dismiss="modal">Close</button><button class="btn btn-primary" type="submit">Save</button></div>
                    </form>
                </div>
            </div>
        </div>
    </body>
</html>`