$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: { //dados que vamos mandar para nossa rota
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){ //201 200 204
        window.location = "/home"
    }).fail(function(erro) { //400 404
        Swal.fire("Ops...", "Erro ao realizar login", "error")
    });
}