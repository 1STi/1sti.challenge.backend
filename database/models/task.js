const mongoose = require('mongoose');

const taskSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true
  },
  status: {
    type: String,
    enum: [
      'A_FAZER',
      'FAZENDO',
      'FEITO'
    ],
    default: 'A_FAZER',
  },
  user: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'User'
  }
}, {
    timestamps: true
  });

module.exports = mongoose.model('Task', taskSchema);