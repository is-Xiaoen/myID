# -*- coding: utf-8 -*-
from flask import Flask, request, jsonify
from paddleocr import PaddleOCR
from PIL import Image
import numpy as np
import base64
import io
import os
import logging

# 设置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

# 初始化PaddleOCR - 使用中文模型
logger.info("Initializing PaddleOCR...")
ocr = PaddleOCR(use_angle_cls=True, lang='ch', use_gpu=False, show_log=False)
logger.info("PaddleOCR initialized successfully")

@app.route('/health', methods=['GET'])
def health():
    return jsonify({"status": "healthy"})

@app.route('/predict/ocr_system', methods=['POST'])
def ocr_predict():
    try:
        # 获取上传的图片
        if 'image' not in request.files:
            return jsonify({"error": "No image file"}), 400
        
        file = request.files['image']
        
        # 保存临时文件
        temp_path = '/tmp/temp_image.jpg'
        file.save(temp_path)
        
        # OCR识别
        result = ocr.ocr(temp_path, cls=True)
        
        # 格式化结果
        formatted_results = []
        if result and len(result) > 0:
            for line in result[0]:
                if line:
                    box = line[0]
                    text = line[1][0]
                    confidence = line[1][1]
                    formatted_results.append([box, [text, confidence]])
        
        # 清理临时文件
        os.remove(temp_path)
        
        return jsonify({
            "results": formatted_results,
            "status": {"code": 0, "message": "success"}
        })
        
    except Exception as e:
        return jsonify({
            "error": str(e),
            "status": {"code": 1, "message": "error"}
        }), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8866)