import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalidades.css';

export default class Personalidades extends Component {
    state = {
        Jogos: []
    }

    componentDidMount() {
        axios.get('http://localhost:8000/api/jogos')
            .then(res => {
                const Jogos = res.data;
                this.setState({ Jogos })
            })
    }

    render() {
        return (
            <div>
                {this.state.Jogos.map((p, id )=>
                    <div className="CardPersonalidades" key={id}>
                        <h1>{p.Nome}</h1>
                        <p>Descrição: {p.Descricao}</p>
                        <p>Preço: R${p.Preco}</p>
                        <p>Nota: {p.Nota}</p>
                    </div>)}
            </div>
        );
    }
}