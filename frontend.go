package main

const tpl = `
<!DOCTYPE html>
<html data-bs-theme="light" lang="en" style>

    <head>
    <meta charset="UTF-8">
    <title>Eric's todo</title>
    <link rel="stylesheet" href="assets/bootstrap/css/bootstrap.min.css">
    <link rel="stylesheet" href="assets/fonts/fontawesome-all.min.css">
    <link rel="stylesheet" href="assets/css/styles.min.css">
    </head>

    <body>
        <section class="vh-100" style="background-color: #e2d5de;">
            <div class="container py-5 h-100">
                <div class="row d-flex justify-content-center align-items-center h-100">
                    <div class="col-xl-10">
                        <div class="card" style="border-radius: 15px;">
                            <div class="card-body p-5">
                                <h6 class="mb-3">Todo lists</h6>
                                <form class="d-flex justify-content-center align-items-center mb-4" action="/add" method="POST">
                                <div class="flex-fill form-outline">
                                    <textarea name="description" id="description" class="form-control form-control-lg" type="text" required placeholder></textarea>
                                    <label class="form-label" for="description">What do you need to do today?</label>
                                </div>
                                    <button class="btn btn-primary btn-lg ms-2" type="submit">Add</button>
                                </form>
                                <ul class="list-group border-0 mb-0">
                                    {{range $i, $a := .}}
                                    <li class="list-group-item border-0 d-flex align-items-center">
                                        <form class="col">
                                            <div class="d-flex form-outline align-items-center">
                                                <input id="form-checkbox-id-0" class="form-check-input me-2" type="checkbox" />
                                                <input class="form-control mb-auto form-control form-control form-label strikethrough" type="text" for="form-checkbox-id-0" value="{{.Description}}" />
                                            </div>
                                        </form>
                                        <a href="#!" data-mdb-toggle="tooltip" title="Remove item"><i class="fas fa-times text-primary"></i></a>
                                    </li>
                                    {{else}}
                                        <div><strong>no rows, click the + button</strong></div>
                                    {{end}}
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </body>
</html>`