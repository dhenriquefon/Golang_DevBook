$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuario");

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem", "error")
        return;
    }

    //fazendo uma requisicao
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: { //dados que vamos mandar para nossa rota
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){ //201 200 204
        Swal.fire("Sucesso...", "Usuário cadastrado com sucesso", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticacar o usuario", "error")            
                })
            })
    }).fail(function(erro) { //400 404
        console.log(erro)
        Swal.fire("Ops...", "Erro ao cadastrar o usuário", "error")
    });
}