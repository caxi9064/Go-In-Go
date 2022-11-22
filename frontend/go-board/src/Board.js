import React from 'react';
import boardpng from './board19x19.png';
import './Board.css';
import Point from "./Point"

const Board = () => {

    let nums = [[]];

    for(let i = 0; i < 19*19; i++){
        nums.push(i)
    }

    return (
        <div class="grid">
            <image src="./board19x19.png"/>
        </div>
    )
}

export default Board;