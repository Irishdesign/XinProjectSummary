CREATE TABLE `orders` (
    `id`             BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `no`             VARCHAR(30)  NOT NULL,
    `type`           TINYINT      NOT NULL ,
    `platform`       VARCHAR(30)  NOT NULL ,
    `member_id`      BIGINT(20)   NOT NULL,
    `status`         TINYINT      NOT NULL,
    `first_name`     VARCHAR(100) NOT NULL,
    `last_name`      VARCHAR(100) NOT NULL,
    `email`          VARCHAR(255) NOT NULL,
    `country`        VARCHAR(100) NOT NULL,
    `address_city`   VARCHAR(50)  NOT NULL,
    `address_area`   VARCHAR(50)  NOT NULL,
    `zip_code`       INT(10)      NULL,
    `address`        VARCHAR(255) NOT NULL,
    `country_code`   VARCHAR(10)  NOT NULL COMMENT '國碼',
    `mobile`         VARCHAR(50)  NOT NULL,
     ...
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`     TIMESTAMP    NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`no`),
    INDEX (`member_id`)
);

CREATE TABLE `order_items`(
    `id`              BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `order_vendor_id` BIGINT(20)   NOT NUll,
    `order_id`        BIGINT(20)   NOT NULL,
     ,,,
    `created_at`      TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`      TIMESTAMP    NULL,
    PRIMARY KEY (`id`),
    INDEX (`order_vendor_id`),
    INDEX (`order_id`),
    INDEX (`product_from`, `product_id`),
    INDEX (`spec_from`, `spec_id`),
    INDEX (`parent_id`)
);

CREATE TABLE `order_vendors` (
    `id`         BIGINT(20) NOT NULL AUTO_INCREMENT,
    `vendor_id`  BIGINT(20) NOT NUll,
    `order_id`   BIGINT(20) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP  NULL,
    PRIMARY KEY (`id`),
    INDEX (`vendor_id`),
    INDEX (`order_id`)
);

CREATE TABLE `order_statuses`(
    `id`         BIGINT(20) NOT NULL AUTO_INCREMENT,
    `order_id`   BIGINT(20) NOT NULL,
    `status`     TINYINT   DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX (`order_id`)
);

CREATE TABLE `cashflows`(
    `id`         BIGINT(20) NOT NULL AUTO_INCREMENT,
    `order_id`   BIGINT(20) NOT NULL,
    `simulation` TINYINT    NOT NULL COMMENT '模擬付款 0=> 是 1=> 否',
    `status`     TINYINT    NOT NULL COMMENT '付款狀態',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `payload`    JSON       NOT NULL,
    PRIMARY KEY (`id`),
    INDEX (`order_id`)
);

CREATE TABLE `vendors` (
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `no`         VARCHAR(30)  NOT NULL,
    `name`       VARCHAR(30)  NULL COMMENT '公司抬頭',
    `bin`        VARCHAR(30)  NULL COMMENT '統一編號',
    `remark`     VARCHAR(255) NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL,
    PRIMARY KEY (`id`)
)